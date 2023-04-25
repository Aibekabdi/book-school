package service

import (
	"book-school/internal/models"
	"book-school/internal/repository"
	hashpassword "book-school/pkg/hash_password"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type SchoolService struct {
	schoolRepo  repository.School
	teacherRepo repository.Teacher
	classRepo   repository.Class
	studentRepo repository.Student
}

func newSchoolService(schoolRepo repository.School, teacherRepo repository.Teacher, classRepo repository.Class, studentRepo repository.Student) *SchoolService {
	return &SchoolService{
		schoolRepo:  schoolRepo,
		teacherRepo: teacherRepo,
		classRepo:   classRepo,
		studentRepo: studentRepo,
	}
}

func (s *SchoolService) Delete(ctx context.Context, schoolId uint) error {
	err := s.schoolRepo.Delete(ctx, schoolId)
	if err != nil {
		return fmt.Errorf("school service: delete: %w", err)
	}
	return nil
}

func (s *SchoolService) Create(ctx context.Context, school models.School) error {
	id, _, err := s.schoolRepo.GetSchool(ctx, school.Name)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	if id != 0 {
		return errors.New("school exist")
	}

	err = validateSchool(school)
	if err != nil {
		return fmt.Errorf("school service: create: %w", err)
	}

	school.Password, err = hashpassword.GenerateHashPassword(school.Password)
	if err != nil {
		return fmt.Errorf("school service: create: %w", err)
	}

	if err := s.schoolRepo.Create(ctx, school); err != nil {
		return fmt.Errorf("school service: create: %w", err)
	}
	return nil
}

func validateSchool(school models.School) error {
	if school.ClassCount <= 0 {
		return fmt.Errorf("validate school: class count is negative")
	}

	if len(school.Name) < 3 {
		return fmt.Errorf("validate school: school name is less or equal 3 symbols")
	}

	if len(school.Password) < 3 {
		return fmt.Errorf("validate school: school password is less or equal 3 symbols")
	}

	return nil
}

func (s *SchoolService) GetById(ctx context.Context, schoolId uint) (models.School, error) {
	ctx = context.WithValue(ctx, models.SchoolId, schoolId)
	school, err := s.schoolRepo.GetAll(ctx)
	if err != nil {
		return models.School{}, fmt.Errorf("school service: get by id: %w", err)
	}
	return school[0], nil
}

func (s *SchoolService) GetAllForSchool(ctx context.Context, schoolId uint) (models.FullSchool, error) {
	var fullSchool models.FullSchool

	school, err := s.schoolRepo.GetAll(context.WithValue(ctx, models.SchoolId, schoolId))
	if err != nil {
		return models.FullSchool{}, fmt.Errorf("school service: get all for school: %w", err)
	}
	fullSchool.School = school[0]

	teachers, err := s.teacherRepo.GetAll(context.WithValue(ctx, models.SchoolId, schoolId))
	if err != nil {
		return models.FullSchool{}, fmt.Errorf("school service: get all for school: %w", err)
	}
	fullSchool.Teachers = make([]models.FullTeacher, len(teachers))

	for i, teacher := range teachers {
		classes, err := s.classRepo.GetAll(context.WithValue(ctx, models.TeacherId, teacher.Id))
		if err != nil {
			return models.FullSchool{}, fmt.Errorf("school service: get all for school: %w", err)
		}
		fullSchool.Teachers[i].Teacher = teacher
		fullSchool.Teachers[i].Classes = make([]models.FullClass, len(classes))
		fullSchool.TotalClasses += len(classes)

		for j, class := range classes {
			fullSchool.Teachers[i].Classes[j].Class = class

			students, err := s.studentRepo.GetAll(context.WithValue(ctx, models.ClassId, class.Id))
			if err != nil {
				return models.FullSchool{}, fmt.Errorf("school service: get all for school: %w", err)
			}

			fullSchool.TotalStudents += len(students)
			fullSchool.Teachers[i].Classes[j].Students = students
		}
	}

	return fullSchool, nil
}

func (s *SchoolService) Update(ctx context.Context, user models.User, update models.SchoolUpdate) error {
	var err error

	update.Password, err = hashpassword.GenerateHashPassword(update.Password)
	if err != nil {
		return err
	}

	if err := s.schoolRepo.Update(ctx, update); err != nil {
		return fmt.Errorf("school service: update: %w", err)
	}
	return nil
}

func (s *SchoolService) GetAllForAdmin(ctx context.Context) ([]models.FullSchool, error) {
	schools, err := s.schoolRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("school service: get all for admin: %w", err)
	}
	fullSchools := make([]models.FullSchool, len(schools))

	for i, school := range schools {
		fullSchools[i], err = s.GetAllForSchool(ctx, school.Id)
		if err != nil {
			return nil, fmt.Errorf("school service: get all for admin: %w", err)
		}
	}

	return fullSchools, nil
}
