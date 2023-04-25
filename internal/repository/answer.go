package repository

import (
	"book-school/internal/models"
	"context"
	"database/sql"
	"fmt"
)

type AnswerRepository struct {
	db *sql.DB
}

func newAnswerRepository(db *sql.DB) *AnswerRepository {
	return &AnswerRepository{
		db: db,
	}
}

func (r *AnswerRepository) GetCompleteAnswers(ctx context.Context, completeTestId uint) ([]models.CompleteAnswers, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE complete_test_id = $1;", completeAnswersTable)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("answer repository: get complete answers: %w", err)
	}
	defer prep.Close()

	var answers []models.CompleteAnswers

	row, err := prep.QueryContext(ctx, completeTestId)
	if err != nil {
		return nil, fmt.Errorf("answer repository: get complete answers: %w", err)
	}
	defer row.Close()

	for row.Next() {
		var answer models.CompleteAnswers
		if err = row.Scan(&answer.Id, &answer.CompleteTestId, &answer.QuestionId, &answer.AnswerId); err != nil {
			return nil, fmt.Errorf("answer repository: get complete answers: %w", err)
		}

		answers = append(answers, answer)
	}

	return answers, nil
}

func (r *AnswerRepository) Create(ctx context.Context, answer models.Answer) error {
	query := fmt.Sprintf("INSERT INTO %s(question_id, with_image, image, audio, answer, correct) VALUES ($1, $2, $3, $4, $5, $6)", answerTable)
	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("answer repository: create: %w", err)
	}
	defer prep.Close()

	if _, err = prep.ExecContext(ctx, answer.QuestionId, answer.WithImage, answer.Image, answer.Audio, answer.Answer, answer.Correct); err != nil {
		return fmt.Errorf("answer repository: create: %w", err)
	}

	return nil
}

func (r *AnswerRepository) GetByQuestionId(ctx context.Context, questionId uint) ([]models.Answer, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE question_id = $1;", answerTable)
	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("answer repository: get by question id: %w", err)
	}
	defer prep.Close()

	rows, err := prep.QueryContext(ctx, questionId)
	if err != nil {
		return nil, fmt.Errorf("answer repository: get by question id: %w", err)
	}
	defer rows.Close()

	var answers []models.Answer

	for rows.Next() {
		var answer models.Answer
		if err = rows.Scan(&answer.Id, &answer.QuestionId, &answer.WithImage, &answer.Image, &answer.Audio, &answer.Answer, &answer.Correct); err != nil {
			return nil, fmt.Errorf("answer repository: get by question id: %w", err)
		}
		answers = append(answers, answer)
	}

	return answers, nil
}

func (r *AnswerRepository) SaveAnswers(ctx context.Context, answer models.CompleteAnswers) error {
	query := fmt.Sprintf("INSERT INTO %s (complete_test_id, question_id, answer_id) VALUES ($1, $2, $3);", completeAnswersTable)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("answer repository: save answers: %w", err)
	}
	defer prep.Close()

	if _, err = prep.ExecContext(ctx, answer.CompleteTestId, answer.QuestionId, answer.AnswerId); err != nil {
		return fmt.Errorf("answer repository: save answers: %w", err)
	}

	return nil
}
