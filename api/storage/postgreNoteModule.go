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

func (s *PostgreNoteModule) Migrate() error {
	exampleNotes := []models.CreateNote{
		models.NewCreateNote(
			1, 1, 1,
			"Teste de nota 1",
		),
		models.NewCreateNote(
			2, 2, 1,
			"Teste de nota 2",
		),
		models.NewCreateNote(
			3, 3, 1,
			"Teste de nota 3",
		),
	}

	for _, n := range exampleNotes {
		if _, err := s.Create(n); err != nil {
			return tracerr.Errorf("failed to create note: %w", tracerr.Wrap(err))
		}
	}

	return nil
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

func (s *PostgreNoteModule) GetFilter(f models.NoteFilter) ([]models.Note, error) {
	var notes []models.Note

	query := s.db.Table("notes")

	if f.Skill != nil {
		query = query.Where("skill_id = ?", *f.Skill)
	}

	if f.Project != nil {
		query = query.Where("project_id = ?", *f.Project)
	}

	if f.User != nil {
		query = query.Where("user_id = ?", *f.User)
	}

	if err := query.Find(&notes).Error; err != nil {
		return nil, tracerr.Errorf("failed to get notes by filter: %w", tracerr.Wrap(err))
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
