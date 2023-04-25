package repository

import (
	"book-school/internal/models"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type TeacherRepository struct {
	db *sql.DB
}

func newTeacherRepository(db *sql.DB) *TeacherRepository {
	return &TeacherRepository{
		db: db,
	}
}

func (r *TeacherRepository) Delete(ctx context.Context, teacherId uint) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", teacherTable)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("teacher repository: delete: %w", err)
	}

	_, err = prep.ExecContext(ctx, teacherId)
	if err != nil {
		return fmt.Errorf("teacher repository: delete: %w", err)
	}

	return nil
}

func (r *TeacherRepository) Update(ctx context.Context, update models.TeacherUpdate) error {
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

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d;", teacherTable, setQuery, argId)
	args = append(args, update.TeacherId)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("teacher repository: update: %w", err)
	}
	defer prep.Close()

	if _, err = prep.ExecContext(ctx, args...); err != nil {
		return fmt.Errorf("teacher repository: update: %w", err)
	}

	return nil
}

func (r *TeacherRepository) Create(ctx context.Context, teacher models.Teacher) error {
	query := fmt.Sprintf("INSERT INTO %s(school_id, first_name, second_name, username, password, private) VALUES ($1, $2, $3, $4, $5, $6)", teacherTable)
	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("teacher repo: create: %w", err)
	}
	defer prep.Close()

	if _, err = prep.ExecContext(ctx, teacher.SchoolId, teacher.FirstName, teacher.SecondName, teacher.Username, teacher.Password, teacher.Private); err != nil {
		return fmt.Errorf("teacher repo: create: %w", err)
	}

	return nil
}

func (r *TeacherRepository) GetTeacher(ctx context.Context, name string) (uint, string, error) {
	query := fmt.Sprintf("SELECT id, password FROM %s WHERE username = $1 LIMIT 1;", teacherTable)
	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, "", fmt.Errorf("teacher repo: get teacher: %w", err)
	}

	var id uint
	var password string

	if err = prep.QueryRowContext(ctx, name).Scan(&id, &password); err != nil {
		return 0, "", fmt.Errorf("teacher repo: get teacher: %w", err)
	}

	return id, password, nil
}

func (r *TeacherRepository) GetById(ctx context.Context, id uint) (models.Teacher, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1;", teacherTable)
	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return models.Teacher{}, fmt.Errorf("teacher repo: get by id: %w", err)
	}
	defer prep.Close()

	var teacher models.Teacher
	if err = prep.QueryRowContext(ctx, id).Scan(&teacher.Id, &teacher.SchoolId, &teacher.FirstName, &teacher.SecondName, &teacher.Username, &teacher.Password, &teacher.Private); err != nil {
		return models.Teacher{}, fmt.Errorf("teacher repo: get by id: %w", err)
	}

	return teacher, nil
}

func (r *TeacherRepository) GetAll(ctx context.Context) ([]models.Teacher, error) {
	args := []interface{}{}

	query := fmt.Sprintf("SELECT * FROM %s;", teacherTable)

	schoolId := ctx.Value(models.SchoolId)
	if schoolId != nil {
		query = strings.ReplaceAll(query, ";", " WHERE school_id = $1;")
		args = append(args, schoolId.(uint))
	}

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("teacher repository: get all: %w", err)
	}
	defer prep.Close()

	var (
		allTeachers []models.Teacher
		oneTeacher  models.Teacher
	)

	rows, err := prep.QueryContext(ctx, args...)
	if err != nil {
		return nil, fmt.Errorf("teacher repository: get all: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&oneTeacher.Id, &oneTeacher.SchoolId, &oneTeacher.FirstName, &oneTeacher.SecondName, &oneTeacher.Username, &oneTeacher.Password, &oneTeacher.Private); err != nil {
			return nil, fmt.Errorf("teacher repository: get all: %w", err)
		}
		allTeachers = append(allTeachers, oneTeacher)
	}

	return allTeachers, nil
}
