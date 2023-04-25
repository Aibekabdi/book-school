package service

import (
	"book-school/internal/models"
	"book-school/internal/repository"
	hashpassword "book-school/pkg/hash_password"
	"context"
	"fmt"
	"strings"
)

type StudentService struct {
	studentRepo repository.Student
	classRepo   repository.Class
	shopRepo    repository.Shop
}

func newStudentService(studentRepo repository.Student, classRepo repository.Class, shopRepo repository.Shop) *StudentService {
	return &StudentService{
		studentRepo: studentRepo,
		classRepo:   classRepo,
		shopRepo:    shopRepo,
	}
}

func (s *StudentService) Delete(ctx context.Context, studentId, teacherId uint) error {
	student, err := s.studentRepo.GetById(ctx, studentId)
	if err != nil {
		return fmt.Errorf("student service: delete: %w", err)
	}

	classes, err := s.classRepo.GetAll(context.WithValue(ctx, models.TeacherId, teacherId))
	if err != nil {
		return fmt.Errorf("student service: delete: %w", err)
	}

	for _, class := range classes {
		if student.ClassId == class.Id {
			err = s.studentRepo.Delete(ctx, student.Id)
			if err != nil {
				return fmt.Errorf("student service: delete: %w", err)
			}
			return nil
		}
	}

	return fmt.Errorf("student service: delete: invalid student")
}

func (s *StudentService) GetStats(ctx context.Context, id uint) (models.Stats, error) {
	student, err := s.studentRepo.GetById(ctx, id)
	if err != nil {
		return models.Stats{}, fmt.Errorf("student service: get stats: %w", err)
	}

	stats, err := s.classRepo.GetStats(context.WithValue(ctx, models.StudentId, id), student.ClassId)
	if err != nil {
		return models.Stats{}, fmt.Errorf("student service: get stats: %w", err)
	}

	stats[0].TotalPoints = stats[0].CreativeTaskPoints + stats[0].AudioPoints + stats[0].TestPoints + stats[0].BookPoints

	return stats[0], nil
}

func (s *StudentService) Create(ctx context.Context, student models.Student, user models.User) error {
	class, err := s.classRepo.GetAll(context.WithValue(ctx, models.ClassId, student.ClassId))
	if err != nil {
		return fmt.Errorf("student service: create: %w", err)
	}

	if len(class) == 0 {
		return fmt.Errorf("student service: create: dont have class")
	}

	if class[0].TeacherId != user.Id {
		return fmt.Errorf("student service: create: not your class")
	}

	err = validateStudent(student)
	if err != nil {
		return fmt.Errorf("student service: create: %w", err)
	}

	student.Password, err = hashpassword.GenerateHashPassword(student.Password)
	if err != nil {
		return fmt.Errorf("student service: create: %w", err)
	}

	id, err := s.studentRepo.Create(ctx, student)
	if err != nil {
		return fmt.Errorf("student service: create: %w", err)
	}

	parts := make(map[string]models.Body)

	for _, bodyPart := range []string{"default head", "default chest", "default legs", "default arms"} {
		parts[strings.Split(bodyPart, " ")[1]], err = s.shopRepo.GetOneBy(context.WithValue(ctx, models.BodyName, bodyPart))
		if err != nil {
			if err = s.studentRepo.Delete(ctx, id); err != nil {
				return fmt.Errorf("student service: create: %w", err)
			}
			return fmt.Errorf("student service: create: %w", err)
		}
	}

	if err = s.studentRepo.SetBody(ctx, id, parts["head"].Id, parts["chest"].Id, parts["legs"].Id, parts["arms"].Id); err != nil {
		if err = s.studentRepo.Delete(ctx, id); err != nil {
			return fmt.Errorf("student service: create: %w", err)
		}
		return fmt.Errorf("student service: create: %w", err)
	}

	for _, bodyId := range []uint{parts["head"].Id, parts["chest"].Id, parts["legs"].Id, parts["arms"].Id} {
		err = s.shopRepo.Buy(ctx, id, bodyId)
		if err != nil {
			if err = s.studentRepo.Delete(ctx, id); err != nil {
				return fmt.Errorf("student service: create: %w", err)
			}
			return fmt.Errorf("student service: create: %w", err)
		}
	}

	return nil
}

func validateStudent(student models.Student) error {
	if len(student.FirstName) < 3 {
		return fmt.Errorf("validate student: first name is less or equal 3 symbols")
	}

	if len(student.SecondName) < 3 {
		return fmt.Errorf("validate student: second name is less or equal 3 symbols")
	}

	if len(student.Username) < 3 {
		return fmt.Errorf("validate student:  username is less or equal 3 symbols")
	}

	if len(student.Password) < 3 {
		return fmt.Errorf("validate student: password is less or equal 3 symbols")
	}

	return nil
}

func (s *StudentService) GetById(ctx context.Context, id uint) (models.Student, error) {
	student, err := s.studentRepo.GetById(ctx, id)
	if err != nil {
		return models.Student{}, fmt.Errorf("student service: get by id: %w", err)
	}
	return student, nil
}

func (s *StudentService) Update(ctx context.Context, user models.User, update models.StudentUpdate) error {
	classes, err := s.classRepo.GetAll(context.WithValue(ctx, models.TeacherId, user.Id))
	if err != nil {
		return fmt.Errorf("student service: create: %w", err)
	}

	student, err := s.studentRepo.GetById(ctx, update.StudentId)
	if err != nil {
		return fmt.Errorf("student service: create: %w", err)
	}

	if !s.isIn(classes, student.ClassId) {
		return fmt.Errorf("student service: create: not your student")
	}

	update.Password, err = hashpassword.GenerateHashPassword(update.Password)
	if err != nil {
		return err
	}

	if err = s.studentRepo.Update(ctx, update); err != nil {
		return fmt.Errorf("student service: create: %w", err)
	}

	return nil
}

func (s *StudentService) isIn(arr []models.Class, classId uint) bool {
	for _, class := range arr {
		if class.Id == classId {
			return true
		}
	}
	return false
}
