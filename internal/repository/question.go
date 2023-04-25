package repository

import (
	"book-school/internal/models"
	"context"
	"database/sql"
	"fmt"
)

type QuestionRepository struct {
	db *sql.DB
}

func newQuestionRepository(db *sql.DB) *QuestionRepository {
	return &QuestionRepository{
		db: db,
	}
}

func (r *QuestionRepository) Create(ctx context.Context, question models.Question) (uint, error) {
	query := fmt.Sprintf("INSERT INTO %s(test_id, with_image, image, audio, question) VALUES ($1, $2, $3, $4, $5) RETURNING id;", questionTable)
	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, fmt.Errorf("question repository: create: %w", err)
	}
	defer prep.Close()

	var id uint
	if err = prep.QueryRowContext(ctx, question.TestId, question.WithImage, question.Image, question.Audio, question.Question).Scan(&id); err != nil {
		return 0, fmt.Errorf("question repository: create: %w", err)
	}

	return id, nil
}

func (r *QuestionRepository) GetByTestId(ctx context.Context, testId uint) ([]models.Question, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE test_id = $1;", questionTable)
	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("question repository: get by test id: %w", err)
	}
	defer prep.Close()

	rows, err := prep.QueryContext(ctx, testId)
	if err != nil {
		return nil, fmt.Errorf("question repository: get by test id: %w", err)
	}
	defer rows.Close()

	var questions []models.Question

	for rows.Next() {
		var question models.Question
		if err = rows.Scan(&question.Id, &question.TestId, &question.WithImage, &question.Image, &question.Audio, &question.Question); err != nil {
			return nil, fmt.Errorf("question repository: get by test id: %w", err)
		}
		questions = append(questions, question)
	}

	return questions, nil
}
