package service

import (
	"book-school/internal/models"
	"book-school/internal/repository"
	hashpassword "book-school/pkg/hash_password"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthService struct {
	schoolRepo  repository.School
	teacherRepo repository.Teacher
	studentRepo repository.Student
	adminRepo   repository.Admin
}

func newAuthService(schoolRepo repository.School, teacherRepo repository.Teacher, studentRepo repository.Student, adminRepo repository.Admin) *AuthService {
	return &AuthService{
		schoolRepo:  schoolRepo,
		teacherRepo: teacherRepo,
		studentRepo: studentRepo,
		adminRepo:   adminRepo,
	}
}

func (s *AuthService) SignIn(ctx context.Context, name, password, role string) (string, error) {
	var (
		id   uint
		pswd string
		err  error
	)
	switch role {
	case models.SchoolRole:
		id, pswd, err = s.schoolRepo.GetSchool(ctx, name)
	case models.TeacherRole:
		id, pswd, err = s.teacherRepo.GetTeacher(ctx, name)
	case models.StudentRole:
		id, pswd, err = s.studentRepo.GetStudent(ctx, name)
	case models.AdminRole:
		id, pswd, err = s.adminRepo.GetAdmin(ctx, name)
	default:
		return "", fmt.Errorf("not supported role")
	}
	if err != nil {
		return "", fmt.Errorf("auth service: sign in: %w", err)
	}

	err = hashpassword.CompareHashAndPassword(pswd, password)
	if err != nil {
		return "", fmt.Errorf("auth service: sign in: %w", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		User: models.User{
			Id:   id,
			Role: role,
		},
	})
	return token.SignedString([]byte(models.Secret))
}

func (s *AuthService) ParseToken(accessToken string) (models.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &models.TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(models.Secret), nil
	})
	if err != nil {
		return models.User{}, err
	}
	claims, ok := token.Claims.(*models.TokenClaims)
	if !ok {
		return models.User{}, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.User, nil
}
