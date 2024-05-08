package storage

import (
	"models"

	"github.com/ztrue/tracerr"
	"gorm.io/gorm"
)

type PostgreNoteModule struct {
	db *gorm.DB
}

func NewPostgreNoteModule(db *gorm.DB) (*PostgreNoteModule, error) {
	return &PostgreNoteModule{db: db}, nil
}

func (s *PostgreNoteModule) Get() ([]models.Note, error) {
	var notes []models.Note
	s.db.Find(&notes)
	return notes, nil
}

func (s *PostgreNoteModule) GetById(id uint64) (*models.Note, error) {
	var note models.Note
	if err := s.db.First(&note, id).Error; err != nil {
		return nil, tracerr.Errorf("failed to get note by id: %w", tracerr.Wrap(err))
	}
	return &note, nil
}

func (s *PostgreNoteModule) GetFilter(f models.NoteFilter) ([]models.NoteDetail, error) {
	var notes []models.NoteDetail

	query := s.db.
		Table("notes").
		Select(`
			notes.id,
			notes.title,
			notes.content,
			users.name AS user,
			skills.name AS skill,
			projects.name AS project
		`).
		Joins("JOIN skills ON notes.skill_id = skills.id").
		Joins("JOIN projects ON notes.project_id = projects.id").
		Joins("JOIN users ON notes.user_id = users.id")

	if f.Skill != nil && *f.Skill != "" {
		query = query.Where("notes.skill_id = ?", *f.Skill)
	}

	if f.Project != nil && *f.Project != "" {
		query = query.Where("notes.project_id = ?", *f.Project)
	}

	if f.User != nil && *f.User != "" {
		query = query.Where("notes.user_id = ?", *f.User)
	}

	result, err := query.Rows()
	if err != nil {
		return nil, tracerr.Errorf("failed to get notes by filter: %w", tracerr.Wrap(err))
	}

	for result.Next() {
		var note models.NoteDetail
		if err := result.Scan(
			&note.Id,
			&note.Title,
			&note.Content,
			&note.User,
			&note.Skill,
			&note.Project,
		); err != nil {
			return nil, tracerr.Errorf("failed to scan note: %w", tracerr.Wrap(err))
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func (s *PostgreNoteModule) GetByProjectId(id uint64) ([]models.Note, error) {
	var notes []models.Note
	s.db.Where("project_id = ?", id).Find(&notes)
	return notes, nil
}

func (s *PostgreNoteModule) GetBySkillId(id uint64) ([]models.Note, error) {
	var notes []models.Note
	s.db.Where("skill_id = ?", id).Find(&notes)
	return notes, nil
}

func (s *PostgreNoteModule) GetByUserId(id uint64) ([]models.Note, error) {
	var notes []models.Note
	s.db.Where("user_id = ?", id).Find(&notes)
	return notes, nil
}

func (s *PostgreNoteModule) Create(n models.CreateNote) (*uint64, error) {
	note := models.Note{
		SkillId:   n.SkillId,
		ProjectId: n.ProjectId,
		UserId:    n.UserId,
		Title:     n.Title,
		Content:   n.Content,
	}
	if err := s.db.Create(&note).Error; err != nil {
		return nil, tracerr.Errorf("failed to create note: %w", tracerr.Wrap(err))
	}
	return &note.Id, nil
}

func (s *PostgreNoteModule) Delete(id uint64) error {
	if err := s.db.Delete(&models.Note{}, id).Error; err != nil {
		return tracerr.Errorf("failed to delete note: %w", tracerr.Wrap(err))
	}
	return nil
}
