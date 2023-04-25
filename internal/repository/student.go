package repository

import (
	"book-school/internal/models"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type StudentRepository struct {
	db *sql.DB
}

func newStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{
		db: db,
	}
}

func (r *StudentRepository) SetBody(ctx context.Context, studentId, headId, chestId, legsId, armsId uint) error {
	query := fmt.Sprintf("INSERT INTO %s(head_id, chest_id, legs_id, arms_id, student_id) VALUES($1, $2, $3, $4, $5);", currentBodyTable)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("student repository: set body: %w", err)
	}
	defer prep.Close()

	_, err = prep.ExecContext(ctx, headId, chestId, legsId, armsId, studentId)
	if err != nil {
		return fmt.Errorf("student repository: set body: %w", err)
	}

	return nil
}

func (r *StudentRepository) UpdateBody(ctx context.Context, studentId, from, to uint, part string) error {
	query := fmt.Sprintf("UPDATE %s SET %s_id = $1 WHERE student_id = $2 AND %s_id = $3", currentBodyTable, part, part)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("student repository: update body: %w", err)
	}
	defer prep.Close()

	_, err = prep.ExecContext(ctx, to, studentId, from)
	if err != nil {
		return fmt.Errorf("student repository: update body: %w", err)
	}

	return nil
}

func (r *StudentRepository) Delete(ctx context.Context, id uint) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", studentTable)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("student repository: delete: %w", err)
	}
	defer prep.Close()

	_, err = prep.ExecContext(ctx, id)
	if err != nil {
		return fmt.Errorf("student repository: delete: %w", err)
	}

	return nil
}

func (r *StudentRepository) Create(ctx context.Context, student models.Student) (uint, error) {
	query := fmt.Sprintf("INSERT INTO %s(class_id, first_name, second_name, username, password) VALUES ($1, $2, $3, $4, $5) RETURNING id", studentTable)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, fmt.Errorf("student repo: create: %w", err)
	}
	defer prep.Close()

	var id uint

	err = prep.QueryRowContext(ctx, student.ClassId, student.FirstName, student.SecondName, student.Username, student.Password).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("student repo: create: %w", err)
	}

	return id, nil
}

func (r *StudentRepository) GetStudent(ctx context.Context, name string) (uint, string, error) {
	query := fmt.Sprintf("SELECT id, password FROM %s WHERE username = $1 LIMIT 1;", studentTable)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, "", fmt.Errorf("student repo: get student: %w", err)
	}

	var id uint
	var password string

	if err = prep.QueryRowContext(ctx, name).Scan(&id, &password); err != nil {
		return 0, "", fmt.Errorf("student repo: get student: %w", err)
	}

	return id, password, nil
}

func (r *StudentRepository) GetById(ctx context.Context, id uint) (models.Student, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1;", studentTable)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return models.Student{}, fmt.Errorf("student repo: get by id: %w", err)
	}
	defer prep.Close()

	var student models.Student
	if err = prep.QueryRowContext(ctx, id).Scan(&student.Id, &student.ClassId, &student.Points, &student.FirstName, &student.SecondName, &student.Username, &student.Password); err != nil {
		return models.Student{}, fmt.Errorf("student repo: get by id: %w", err)
	}

	return student, nil
}

func (r *StudentRepository) GetAll(ctx context.Context) ([]models.Student, error) {
	args := []interface{}{}

	query := fmt.Sprintf("SELECT * FROM %s;", studentTable)

	classId := ctx.Value(models.ClassId)
	if classId != nil {
		query = strings.ReplaceAll(query, ";", " WHERE class_id = $1;")
		args = append(args, classId.(uint))
	}

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("student repo: get all: %w", err)
	}
	defer prep.Close()

	row, err := prep.QueryContext(ctx, args...)
	if err != nil {
		return nil, fmt.Errorf("student repo: get all: %w", err)
	}
	defer row.Close()

	var (
		students []models.Student
		student  models.Student
	)

	for row.Next() {
		if err = row.Scan(&student.Id, &student.ClassId, &student.Points, &student.FirstName, &student.SecondName, &student.Username, &student.Password); err != nil {
			return nil, fmt.Errorf("student repo: get all: %w", err)
		}

		students = append(students, student)
	}

	return students, nil
}

func (r *StudentRepository) Update(ctx context.Context, update models.StudentUpdate) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if update.FirstName != "" {
		setValues = append(setValues, fmt.Sprintf("first_name = $%d", argId))
		args = append(args, update.FirstName)
		argId++
	}

	if update.SecondName != "" {
		setValues = append(setValues, fmt.Sprintf("second_name = $%d", argId))
		args = append(args, update.SecondName)
		argId++
	}

	if update.Username != "" {
		setValues = append(setValues, fmt.Sprintf("username = $%d", argId))
		args = append(args, update.Username)
		argId++
	}

	if update.Password != "" {
		setValues = append(setValues, fmt.Sprintf("password = $%d", argId))
		args = append(args, update.Password)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d;", studentTable, setQuery, argId)
	args = append(args, update.StudentId)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("student repository: update: %w", err)
	}
	defer prep.Close()

	if _, err = prep.ExecContext(ctx, args...); err != nil {
		return fmt.Errorf("student repository: update: %w", err)
	}

	return nil
}

func (r *StudentRepository) GivePoints(ctx context.Context, studentId, points uint) error {
	query := fmt.Sprintf("UPDATE %s SET points = points + $1 WHERE id = $2;", studentTable)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("student repository: give points: %w", err)
	}
	defer prep.Close()

	if _, err = prep.ExecContext(ctx, points, studentId); err != nil {
		return fmt.Errorf("student repository: give points: %w", err)
	}

	return nil
}

func (r *StudentRepository) TakePoints(ctx context.Context, studentId, points uint) error {
	query := fmt.Sprintf("UPDATE %s SET points = points - $1 WHERE id = $2;", studentTable)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("student repository: take points: %w", err)
	}
	defer prep.Close()

	if _, err = prep.ExecContext(ctx, points, studentId); err != nil {
		return fmt.Errorf("student repository: take points: %w", err)
	}

	return nil
}
