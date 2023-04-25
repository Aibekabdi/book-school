package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type AudioRepository struct {
	db *sql.DB
}

func newAudioRepository(db *sql.DB) *AudioRepository {
	return &AudioRepository{
		db: db,
	}
}

func (r *AudioRepository) CheckCompleteAudio(ctx context.Context, bookId, studentId uint) error {
	query := fmt.Sprintf("SELECT id FROM %s WHERE book_id = $1 AND student_id = $2", completeAudioTable)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("audio repository: check complete audio: %w", err)
	}
	defer prep.Close()

	var id uint

	err = prep.QueryRowContext(ctx, bookId, studentId).Scan(&id)
	if err != nil {
		return fmt.Errorf("audio repository: check complete audio: %w", err)
	}

	return nil
}

func (r *AudioRepository) Complete(ctx context.Context, bookId, studentId, points uint) error {
	query := fmt.Sprintf("INSERT INTO %s(book_id, student_id, points) VALUES($1, $2, $3);", completeAudioTable)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("audio repository: complete: %w", err)
	}
	defer prep.Close()

	if _, err = prep.ExecContext(ctx, bookId, studentId, points); err != nil {
		return fmt.Errorf("audio repository: complete: %w", err)
	}

	return nil
}
