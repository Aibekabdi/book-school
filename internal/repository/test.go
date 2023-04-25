package repository

import (
	"book-school/internal/models"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type TestRepository struct {
	db *sql.DB
}

func newTestRepository(db *sql.DB) *TestRepository {
	return &TestRepository{
		db: db,
	}
}

func (r *TestRepository) Create(ctx context.Context, bookId uint, lang string) (uint, error) {
	query := fmt.Sprintf("INSERT INTO %s(book_id, lang) VALUES ($1, $2) RETURNING id;", testTable)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, fmt.Errorf("test repository: create: %w", err)
	}
	defer prep.Close()

	var id uint
	if err = prep.QueryRowContext(ctx, bookId, lang).Scan(&id); err != nil {
		return 0, fmt.Errorf("test repository: create: %w", err)
	}

	return id, nil
}

func (r *TestRepository) Get(ctx context.Context) ([]models.Test, error) {
	args := []interface{}{}
	argsStr := []string{}
	argsNum := 1

	bookId := ctx.Value(models.BookId)
	if bookId != nil {
		argsStr = append(argsStr, fmt.Sprintf("book_id = $%d", argsNum))
		args = append(args, bookId.(uint))
		argsNum++
	}

	testId := ctx.Value(models.TestId)
	if testId != nil {
		argsStr = append(argsStr, fmt.Sprintf("id = $%d", argsNum))
		args = append(args, testId.(uint))
		argsNum++
	}

	whereCondition := ""
	if len(argsStr) != 0 {
		whereCondition = " WHERE " + strings.Join(argsStr, " AND ")
	}

	query := fmt.Sprintf("SELECT * FROM %s%s;", testTable, whereCondition)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("test repository: get by book id: %w", err)
	}
	defer prep.Close()

	row, err := prep.QueryContext(ctx, args...)
	if err != nil {
		return nil, fmt.Errorf("test repository: get by book id: %w", err)
	}

	var tests []models.Test

	for row.Next() {
		var test models.Test
		if err = row.Scan(&test.Id, &test.BookId, &test.Lang); err != nil {
			return nil, fmt.Errorf("test repository: get by book id: %w", err)
		}

		tests = append(tests, test)
	}

	return tests, nil
}

func (r *TestRepository) Delete(ctx context.Context, testId uint) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1;", testTable)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("test repository: delete: %w", err)
	}
	defer prep.Close()

	if _, err = prep.ExecContext(ctx, testId); err != nil {
		return fmt.Errorf("test repository: delete: %w", err)
	}

	return nil
}

func (r *TestRepository) SaveTest(ctx context.Context, test models.CompleteTest) (uint, error) {
	query := fmt.Sprintf("INSERT INTO %s (student_id, test_id, points) VALUES ($1, $2, $3) RETURNING id;", completeTestsTable)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, fmt.Errorf("test repository: save test: %w", err)
	}
	defer prep.Close()

	var id uint
	if err = prep.QueryRowContext(ctx, test.StudentId, test.TestId, test.Points).Scan(&id); err != nil {
		return 0, fmt.Errorf("test repository: save test: %w", err)
	}

	return id, nil
}

func (r TestRepository) DeleteSavedTest(ctx context.Context, id uint) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1;", completeTestsTable)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("test repository: delete: %w", err)
	}
	defer prep.Close()

	if _, err = prep.ExecContext(ctx, id); err != nil {
		return fmt.Errorf("test repository: delete: %w", err)
	}

	return nil
}

func (r *TestRepository) GetCompleteTest(ctx context.Context) (models.CompleteTest, error) {
	args := []interface{}{}
	argsStr := []string{}
	argsNum := 1

	whereOpt := ""

	testIdCtx := ctx.Value(models.TestId)
	if testIdCtx != nil {
		argsStr = append(argsStr, fmt.Sprintf("test_id = $%d", argsNum))
		args = append(args, testIdCtx.(uint))
		argsNum++
	}

	studentIdCtx := ctx.Value(models.StudentId)
	if studentIdCtx != nil {
		argsStr = append(argsStr, fmt.Sprintf("student_id = $%d", argsNum))
		args = append(args, studentIdCtx.(uint))
		argsNum++
	}

	completeTestIdCtx := ctx.Value(models.CompleteTestId)
	if completeTestIdCtx != nil {
		argsStr = append(argsStr, fmt.Sprintf("id = $%d", argsNum))
		args = append(args, completeTestIdCtx.(uint))
		argsNum++
	}

	if len(argsStr) != 0 {
		whereOpt = "WHERE " + strings.Join(argsStr, " AND ")
	}

	// WHERE test_id = $1 AND student_id = $2;

	query := fmt.Sprintf("SELECT * FROM %s %s", completeTestsTable, whereOpt)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return models.CompleteTest{}, fmt.Errorf("test repository: get complete test: %w", err)
	}
	defer prep.Close()

	var test models.CompleteTest

	if err = prep.QueryRowContext(ctx, args...).Scan(&test.Id, &test.TestId, &test.StudentId, &test.Points); err != nil {
		return models.CompleteTest{}, fmt.Errorf("test repository: get complete test: %w", err)
	}

	return test, nil
}
