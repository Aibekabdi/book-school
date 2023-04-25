package service

import (
	"book-school/internal/models"
	"book-school/internal/repository"
	"context"
	"fmt"
	"sort"
)

type ClassService struct {
	classRepo   repository.Class
	studentRepo repository.Student
	teacherRepo repository.Teacher
	schoolRepo  repository.School
}

func newClassService(classRepo repository.Class, studentRepo repository.Student, teacherRepo repository.Teacher, schoolRepo repository.School) *ClassService {
	return &ClassService{
		classRepo:   classRepo,
		studentRepo: studentRepo,
		teacherRepo: teacherRepo,
		schoolRepo:  schoolRepo,
	}
}

func (r *ClassService) Delete(ctx context.Context, classId uint, user models.User) error {
	c, err := r.classRepo.GetAll(context.WithValue(ctx, models.ClassId, classId))
	if err != nil {
		return fmt.Errorf("class service: delete: %w", err)
	}

	if user.Role == models.TeacherRole {
		teacher, err := r.teacherRepo.GetById(ctx, user.Id)
		if err != nil {
			return fmt.Errorf("class service: delete: %w", err)
		}

		if !teacher.Private {
			return fmt.Errorf("class service: delete: not private teacher")
		}

		if len(c) == 0 || c[0].TeacherId != user.Id {
			return fmt.Errorf("class service: delete: invalid class")
		}

		return nil
	}

	if len(c) == 0 || c[0].SchoolId != user.Id {
		return fmt.Errorf("class service: delete: invalid class")
	}

	err = r.classRepo.Delete(ctx, classId)
	if err != nil {
		return fmt.Errorf("class service: delete: %w", err)
	}

	return nil
}

func (r *ClassService) Create(ctx context.Context, class models.Class, user models.User) error {
	if user.Role == models.SchoolRole {
		school, err := r.schoolRepo.GetAll(context.WithValue(ctx, models.SchoolId, user.Id))
		if err != nil {
			return fmt.Errorf("class service: create: %w", err)
		}

		teacher, err := r.teacherRepo.GetById(ctx, class.TeacherId)
		if err != nil {
			return fmt.Errorf("class service: create: %w", err)
		}

		classes, err := r.classRepo.GetAll(context.WithValue(ctx, models.SchoolId, user.Id))
		if err != nil {
			return fmt.Errorf("class service: create: %w", err)
		}

		if int(school[0].ClassCount) == len(classes) {
			return fmt.Errorf("class service: create: school limit on class create")
		}

		if teacher.SchoolId != user.Id {
			return fmt.Errorf("class service: create: not your teacher")
		}

		_, err = r.classRepo.GetClass(ctx, class.Name, class.Grade, class.SchoolId)
		if err == nil {
			return fmt.Errorf("class service: create: class exists")
		}

		class.SchoolId = user.Id
	} else if user.Role == models.TeacherRole {
		teacher, err := r.teacherRepo.GetById(ctx, user.Id)
		if err != nil {
			return fmt.Errorf("class service: create: %w", err)
		}

		classes, err := r.classRepo.GetAll(context.WithValue(ctx, models.TeacherId, user.Id))
		if err != nil {
			return fmt.Errorf("class service: create: %w", err)
		}

		if len(classes) == 3 {
			return fmt.Errorf("class service: create: private teacher limit on class create")
		}

		if !teacher.Private {
			return fmt.Errorf("class service: create: you cant create class")
		}

		_, err = r.classRepo.GetClass(ctx, class.Name, class.Grade, 1)
		if err == nil {
			return fmt.Errorf("class service: create: class exists")
		}

		class.SchoolId = 1
		class.TeacherId = user.Id
	}

	if class.Grade != "2 год" && class.Grade != "3 год" && class.Grade != "4 год" && class.Grade != "5 год" &&
		class.Grade != "1 класс" && class.Grade != "2 класс" && class.Grade != "3 класс" && class.Grade != "4 класс" {
		return fmt.Errorf("class service: create: invalid class grade")
	}

	if err := r.classRepo.Create(ctx, class); err != nil {
		return fmt.Errorf("class service: create: %w", err)
	}

	return nil
}

func (r *ClassService) GetStatsTotal(ctx context.Context, user models.User) ([]models.Stats, error) {
	stats, err := r.GetStats(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("class service: get stats total: %w", err)
	}

	var info []models.Stats

	for _, stat := range stats {
		info = append(info, stat.Stats...)
	}

	sort.Slice(info, func(i, j int) bool {
		return info[i].TotalPoints > info[j].TotalPoints
	})

	return info, nil
}

func (r *ClassService) GetStats(ctx context.Context, user models.User) ([]models.ClassStats, error) {
	var (
		stats   []models.ClassStats
		classes []models.Class
		err     error
	)

	if user.Role == models.TeacherRole {
		classes, err = r.classRepo.GetAll(context.WithValue(ctx, models.TeacherId, user.Id))
	} else if user.Role == models.SchoolRole {
		classes, err = r.classRepo.GetAll(context.WithValue(ctx, models.SchoolId, user.Id))
	}
	if err != nil {
		return nil, fmt.Errorf("class service: get stats: %w", err)
	}

	for _, class := range classes {
		stat, err := r.classRepo.GetStats(ctx, class.Id)
		if err != nil {
			return nil, fmt.Errorf("class service: get stats: %w", err)
		}

		for i, s := range stat {
			stat[i].TotalPoints = s.AudioPoints + s.TestPoints + s.BookPoints + s.CreativeTaskPoints
		}

		sort.Slice(stat, func(i, j int) bool {
			return stat[i].TotalPoints > stat[j].TotalPoints
		})

		stats = append(stats, models.ClassStats{Class: class, Stats: stat})
	}

	return stats, nil
}
