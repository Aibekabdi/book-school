package service

import (
	"book-school/internal/models"
	"book-school/internal/repository"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type AudioService struct {
	audioRepo   repository.Audio
	studentRepo repository.Student
}

func newAudioService(audioRepo repository.Audio, studentRepo repository.Student) *AudioService {
	return &AudioService{
		audioRepo:   audioRepo,
		studentRepo: studentRepo,
	}
}

func (s *AudioService) Complete(ctx context.Context, bookId, studentId uint) (uint, error) {
	err := s.audioRepo.CheckCompleteAudio(ctx, bookId, studentId)
	if err == nil {
		return 0, fmt.Errorf("audio service: complete: you passed audio once")
	} else if !errors.Is(err, sql.ErrNoRows) {
		return 0, fmt.Errorf("audio service: complete: %w", err)
	}

	if err := s.studentRepo.GivePoints(ctx, studentId, models.AudioPoints); err != nil {
		return 0, fmt.Errorf("audio service: complete: %w", err)
	}

	if err := s.audioRepo.Complete(ctx, bookId, studentId, models.AudioPoints); err != nil {
		if err := s.studentRepo.TakePoints(ctx, studentId, models.AudioPoints); err != nil {
			return 0, fmt.Errorf("book service: complete: %w", err)
		}
		return 0, fmt.Errorf("audio service: complete: %w", err)
	}

	return models.AudioPoints, nil
}
