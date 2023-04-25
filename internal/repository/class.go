package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"book-school/internal/models"
)

type ClassRepository struct {
	db *sql.DB
}

func newClassRepository(db *sql.DB) *ClassRepository {
	return &ClassRepository{
		db: db,
	}
}

func (r *ClassRepository) Delete(ctx context.Context, classId uint) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1;", classTable)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("class repository: delete: %w", err)
	}
	defer prep.Close()

	_, err = prep.ExecContext(ctx, classId)
	if err != nil {
		return fmt.Errorf("class repository: delete: %w", err)
	}

	return nil
}

func (r *ClassRepository) GetAll(ctx context.Context) ([]models.Class, error) {
	args := []interface{}{}

	query := fmt.Sprintf("SELECT * FROM %s;", classTable)

	schoolId := ctx.Value(models.SchoolId)
	if schoolId != nil {
		query = strings.ReplaceAll(query, ";", " WHERE school_id = $1;")
		args = append(args, schoolId.(uint))
	}

	teacherId := ctx.Value(models.TeacherId)
	if teacherId != nil {
		query = strings.ReplaceAll(query, ";", " WHERE teacher_id = $1;")
		args = append(args, teacherId.(uint))
	}

	id := ctx.Value(models.ClassId)
	if id != nil {
		query = strings.ReplaceAll(query, ";", " WHERE id = $1;")
		args = append(args, id.(uint))
	}

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("class repository: get all: %w", err)
	}
	defer prep.Close()

	var (
		classes []models.Class
		class   models.Class
	)

	row, err := prep.QueryContext(ctx, args...)
	if err != nil {
		return nil, fmt.Errorf("class repository: get all: %w", err)
	}
	defer row.Close()

	for row.Next() {
		if err = row.Scan(&class.Id, &class.SchoolId, &class.TeacherId, &class.Grade, &class.Name); err != nil {
			return nil, fmt.Errorf("class repository: get all: %w", err)
		}

		classes = append(classes, class)
	}

	return classes, nil
}

func (r *ClassRepository) GetClass(ctx context.Context, name, grade string, schoolId uint) (models.Class, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE name = $1 AND grade = $2 AND school_id = $3;", classTable)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return models.Class{}, fmt.Errorf("class service: get class: %w", err)
	}
	defer prep.Close()

	var class models.Class

	if err = prep.QueryRowContext(ctx, name, grade, schoolId).Scan(&class.Id, &class.SchoolId, &class.TeacherId, &class.Grade, &class.Name); err != nil {
		return models.Class{}, fmt.Errorf("class service: get class: %w", err)
	}

	return class, nil
}

func (r *ClassRepository) Create(ctx context.Context, class models.Class) error {
	query := fmt.Sprintf("INSERT INTO %s(school_id, teacher_id, grade, name) VALUES($1, $2, $3, $4);", classTable)
	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("class repository: create: %w", err)
	}
	defer prep.Close()

	if _, err = prep.ExecContext(ctx, class.SchoolId, class.TeacherId, class.Grade, class.Name); err != nil {
		return fmt.Errorf("class repository: create: %w", err)
	}
	return nil
}

/*
	SELECT
		s.id,
		s.first_name,
		s.second_name,
		c.grade,
		c.name,
		COALESCE((select sum(b.points) from complete_books b where b.student_id = s.id), 0) as book_points,
		COALESCE((select sum(a.points) from complete_audios a where a.student_id = s.id), 0) as audio_points,
		COALESCE((select sum(t.points) from complete_tests t where t.student_id = s.id), 0) as test_points,
		COALESCE((select sum(o.points) from open_comments o where o.student_id = s.id), 0) as creative_task_points
	FROM students s
	INNER JOIN classes c ON c.id = s.class_id
	WHERE s.class_id = 1;
*/

func (r *ClassRepository) GetStats(ctx context.Context, classId uint) ([]models.Stats, error) {
	args := []interface{}{}

	args = append(args, classId)

	student := ""

	studentId := ctx.Value(models.StudentId)
	if studentId != nil {
		student = " AND s.id = $2"
		args = append(args, studentId.(uint))
	}

	/*

	 */

	query := fmt.Sprintf(`
	SELECT
		s.id,
		s.first_name,
		s.second_name,
		c.grade,
		c.name,
		COALESCE((select sum(b.points) from %s b where b.student_id = s.id), 0) as book_points,
		COALESCE((select sum(a.points) from %s a where a.student_id = s.id), 0) as audio_points,
		COALESCE((select sum(t.points) from %s t where t.student_id = s.id), 0) as test_points,
		COALESCE((select sum(c.points) from %s c where c.student_id = s.id), 0) as creative_task_points,
		COALESCE((select sum(o.points) from %s o where o.student_id = s.id), 0) as open_points
	FROM %s s
	INNER JOIN %s c ON c.id = s.class_id
	WHERE s.class_id = $1%s;
	`, completeBooksTable, completeAudioTable, completeTestsTable, "creative_comments", "open_comments", studentTable, classTable, student)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("class repository: get stats: %w", err)
	}
	defer prep.Close()

	rows, err := prep.QueryContext(ctx, args...)
	if err != nil {
		return nil, fmt.Errorf("class repository: get stats: %w", err)
	}

	var stats []models.Stats

	for rows.Next() {
		var stat models.Stats
		if err := rows.Scan(&stat.StudentId, &stat.StudentFirstName, &stat.StudentSecondName, &stat.Grade, &stat.Name, &stat.BookPoints, &stat.AudioPoints, &stat.TestPoints, &stat.CreativeTaskPoints, &stat.OpenPoints); err != nil {
			return nil, fmt.Errorf("class repository: get stats: %w", err)
		}

		stats = append(stats, stat)
	}

	return stats, nil
}
