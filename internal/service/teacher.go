package service

import (
	"book-school/internal/models"
	"book-school/internal/repository"
	hashpassword "book-school/pkg/hash_password"
	"context"
	"fmt"
)

type TeacherService struct {
	teacherRepo repository.Teacher
	classRepo   repository.Class
	studentRepo repository.Student
	schoolRepo  repository.School
}

func newTeacherService(teacherRepo repository.Teacher, classRepo repository.Class, studentRepo repository.Student, schoolRepo repository.School) *TeacherService {
	return &TeacherService{
		teacherRepo: teacherRepo,
		classRepo:   classRepo,
		studentRepo: studentRepo,
		schoolRepo:  schoolRepo,
	}
}

func (s *TeacherService) Delete(ctx context.Context, teacherId, schoolId uint) error {
	teacher, err := s.teacherRepo.GetById(ctx, teacherId)
	if err != nil {
		return fmt.Errorf("teacher service: delete: %w", err)
	}

	if teacher.SchoolId != schoolId {
		return fmt.Errorf("teacher service: delete: not your teacher")
	}

	err = s.teacherRepo.Delete(ctx, teacherId)
	if err != nil {
		return fmt.Errorf("teacher service: delete: %w", err)
	}

	return nil
}

func (s *TeacherService) Create(ctx context.Context, teacher models.Teacher, private bool, schoolId uint) error {
	err := validateTeacher(teacher)
	if err != nil {
		return fmt.Errorf("teacher service: create: %w", err)
	}

	teacher.Password, err = hashpassword.GenerateHashPassword(teacher.Password)
	if err != nil {
		return fmt.Errorf("teacher service: create: %w", err)
	}

	teacher.SchoolId = schoolId
	teacher.Private = private
	if err := s.teacherRepo.Create(ctx, teacher); err != nil {
		return fmt.Errorf("teacher service: create: %w", err)
	}
	return nil
}

func validateTeacher(teacher models.Teacher) error {
	if len(teacher.FirstName) < 3 {
		return fmt.Errorf("validate teacher: first name is less or equal 3 symbols")
	}

	if len(teacher.SecondName) < 3 {
		return fmt.Errorf("validate teacher: second name is less or equal 3 symbols")
	}

	if len(teacher.Username) < 3 {
		return fmt.Errorf("validate teacher:  username is less or equal 3 symbols")
	}

	if len(teacher.Password) < 3 {
		return fmt.Errorf("validate teacher: password is less or equal 3 symbols")
	}

	return nil
}

func (s *TeacherService) GetById(ctx context.Context, id uint) (models.Teacher, error) {
	teacher, err := s.teacherRepo.GetById(ctx, id)
	if err != nil {
		return models.Teacher{}, fmt.Errorf("teacher service: get by id: %w", err)
	}
	return teacher, nil
}

func (s *TeacherService) GetAllForTeacher(ctx context.Context, teacherId uint) (models.FullTeacher, error) {
	var fullTeacher models.FullTeacher

	teacher, err := s.teacherRepo.GetById(ctx, teacherId)
	if err != nil {
		return models.FullTeacher{}, fmt.Errorf("teacher service: get all for teacher: %w", err)
	}
	fullTeacher.Teacher = teacher

	school, err := s.schoolRepo.GetAll(context.WithValue(ctx, models.SchoolId, teacher.SchoolId))
	if err != nil {
		return models.FullTeacher{}, fmt.Errorf("teacher service: get all for teacher: %w", err)
	}
	fullTeacher.School = school[0]

	classes, err := s.classRepo.GetAll(context.WithValue(ctx, models.TeacherId, teacher.Id))
	if err != nil {
		return models.FullTeacher{}, fmt.Errorf("teacher service: get all for teacher: %w", err)
	}
	fullTeacher.Classes = make([]models.FullClass, len(classes))

	for i, class := range classes {
		fullTeacher.Classes[i].Class = class

		students, err := s.studentRepo.GetAll(context.WithValue(ctx, models.ClassId, class.Id))
		if err != nil {
			return models.FullTeacher{}, fmt.Errorf("teacher service: get all for teacher: %w", err)
		}

		fullTeacher.Classes[i].Students = students
	}

	return fullTeacher, nil
}

func (s *TeacherService) Update(ctx context.Context, user models.User, update models.TeacherUpdate) error {
	if user.Role == models.SchoolRole {
		teachers, err := s.teacherRepo.GetAll(context.WithValue(ctx, models.SchoolId, user.Id))
		if err != nil {
			return fmt.Errorf("teacher service: update: %w", err)
		}

		if !s.isIn(teachers, update.TeacherId) {
			return fmt.Errorf("teacher service: update: not your teacher")
		}
	}

	hash, err := hashpassword.GenerateHashPassword(update.Password)
	if err != nil {
		return err
	}

	update.Password = hash

	if err := s.teacherRepo.Update(ctx, update); err != nil {
		return fmt.Errorf("teacher service: update: %w", err)
	}
	return nil
}

func (s *TeacherService) isIn(arr []models.Teacher, teacherId uint) bool {
	for _, teacher := range arr {
		if teacher.Id == teacherId {
			return true
		}
	}
	return false
}
