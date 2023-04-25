package repository

import (
	"book-school/internal/models"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type SchoolRepository struct {
	db *sql.DB
}

func newSchoolRepository(db *sql.DB) *SchoolRepository {
	return &SchoolRepository{
		db: db,
	}
}

func (r *SchoolRepository) Delete(ctx context.Context, schoolId uint) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1;", schoolTable)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("school repository: delete: %w", err)
	}
	defer prep.Close()

	_, err = prep.ExecContext(ctx, schoolId)
	if err != nil {
		return fmt.Errorf("school repository: delete: %w", err)
	}

	return nil
}

func (r *SchoolRepository) Update(ctx context.Context, update models.SchoolUpdate) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if update.ClassCount > 0 {
		setValues = append(setValues, fmt.Sprintf("class_count = $%d", argId))
		args = append(args, update.ClassCount)
		argId++
	}

	if update.Name != "" {
		setValues = append(setValues, fmt.Sprintf("name = $%d", argId))
		args = append(args, update.Name)
		argId++
	}

	if update.Password != "" {
		setValues = append(setValues, fmt.Sprintf("password = $%d", argId))
		args = append(args, update.Password)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d;", schoolTable, setQuery, argId)
	args = append(args, update.SchoolId)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("school repository: update: %w", err)
	}
	defer prep.Close()

	if _, err = prep.ExecContext(ctx, args...); err != nil {
		return fmt.Errorf("school repository: update: %w", err)
	}

	return nil
}

func (r *SchoolRepository) Create(ctx context.Context, school models.School) error {
	query := fmt.Sprintf("INSERT INTO %s(class_count, name, password) VALUES($1, $2, $3);", schoolTable)
	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("school repo: create: %w", err)
	}

	if _, err = prep.ExecContext(ctx, school.ClassCount, school.Name, school.Password); err != nil {
		return fmt.Errorf("school repo: create: %w", err)
	}

	return nil
}

func (r *SchoolRepository) GetSchool(ctx context.Context, name string) (uint, string, error) {
	query := fmt.Sprintf("SELECT id, password FROM %s WHERE name = $1 LIMIT 1", schoolTable)
	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, "", fmt.Errorf("school repo: get school: %w", err)
	}

	var id uint
	var password string

	if err = prep.QueryRowContext(ctx, name).Scan(&id, &password); err != nil {
		return 0, "", fmt.Errorf("school repo: get school: %w", err)
	}

	return id, password, nil
}

func (r *SchoolRepository) GetAll(ctx context.Context) ([]models.School, error) {
	args := []interface{}{}

	query := fmt.Sprintf("SELECT * FROM %s;", schoolTable)

	id := ctx.Value(models.SchoolId)
	if id != nil {
		query = strings.ReplaceAll(query, ";", " WHERE id = $1;")
		args = append(args, id.(uint))
	}

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("school repo: get by id: %w", err)
	}

	var (
		schools []models.School
		school  models.School
	)

	row, err := prep.QueryContext(ctx, args...)
	if err != nil {
		return nil, fmt.Errorf("school repo: get by id: %w", err)
	}
	defer row.Close()

	for row.Next() {
		if err = row.Scan(&school.Id, &school.ClassCount, &school.Name, &school.Password); err != nil {
			return nil, fmt.Errorf("school repo: get by id: %w", err)
		}

		schools = append(schools, school)
	}

	return schools, nil
}
