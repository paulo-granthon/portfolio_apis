package service

import (
	"storage"

	"github.com/ztrue/tracerr"
)

type Service struct {
	NoteService *NoteService
}

func NewService(s storage.Storage) (*Service, error) {
	noteService, err := NewNoteService(s)
	if err != nil {
		return nil, tracerr.Errorf("failed to create note service: %v", err)
	}

	return &Service{
		NoteService: noteService,
	}, nil
}
