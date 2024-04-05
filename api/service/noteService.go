package service

import (
	"models"
	"storage"

	"github.com/ztrue/tracerr"
)

type NoteService struct {
	noteStorage    *storage.NoteStorageModule
	skillStorage   *storage.SkillStorageModule
	projectStorage *storage.ProjectStorageModule
	userStorage    *storage.UserStorageModule
}

func NewNoteService(
	storage storage.Storage,
) (*NoteService, error) {
	notestorage, err := storage.GetNoteModule()
	if err != nil {
		return nil, tracerr.Errorf("failed to create note service: failed to get NoteModule from storage: %w", tracerr.Wrap(err))
	}

	skillStorage, err := storage.GetSkillModule()
	if err != nil {
		return nil, tracerr.Errorf("failed to create note service: failed to get SkillModule from storage: %w", tracerr.Wrap(err))
	}

	projectStorage, err := storage.GetProjectModule()
	if err != nil {
		return nil, tracerr.Errorf("failed to create note service: failed to get ProjectModule from storage: %w", tracerr.Wrap(err))
	}

	userStorage, err := storage.GetUserModule()
	if err != nil {
		return nil, tracerr.Errorf("failed to create note service: failed to get UserModule from storage: %w", tracerr.Wrap(err))
	}

	return &NoteService{
		noteStorage:    &notestorage,
		skillStorage:   &skillStorage,
		projectStorage: &projectStorage,
		userStorage:    &userStorage,
	}, nil
}

func (s *NoteService) Create(n models.CreateNoteByNames) (*uint64, error) {
	skillStorage := *s.skillStorage
	skill, err := skillStorage.GetByName(n.Skill)
	if err != nil {
		return nil, tracerr.Errorf("failed to create note: failed to get skill by name: %w", tracerr.Wrap(err))
	}

	projectStorage := *s.projectStorage
	project, err := projectStorage.GetByName(n.Project)
	if err != nil {
		return nil, tracerr.Errorf("failed to create note: failed to get project by name: %w", tracerr.Wrap(err))
	}

	userStorage := *s.userStorage
	user, err := userStorage.GetByName(n.User)
	if err != nil {
		return nil, tracerr.Errorf("failed to create note: failed to get user by name: %w", tracerr.Wrap(err))
	}

	note := models.CreateNote{
		SkillId:   skill.Id,
		ProjectId: project.Id,
		UserId:    user.Id,
		Content:   n.Content,
	}

	noteStorage := *s.noteStorage

	id, err := noteStorage.Create(note)
	if err != nil {
		return nil, tracerr.Errorf("failed to create note: %w", tracerr.Wrap(err))
	}

	return id, nil
}
