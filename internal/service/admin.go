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

type AdminService struct {
	adminRepo repository.Admin
}

func newAdminService(adminRepo repository.Admin) *AdminService {
	return &AdminService{
		adminRepo: adminRepo,
	}
}

func (s *AdminService) Create(ctx context.Context, admin models.Admin) error {
	id, _, err := s.adminRepo.GetAdmin(ctx, admin.Username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	if id != 0 {
		return fmt.Errorf("admin exist")
	}

	admin.Password, err = hashpassword.GenerateHashPassword(admin.Password)
	if err != nil {
		return fmt.Errorf("admin service: create: %w")
	}

	err = s.adminRepo.Create(ctx, admin)
	if err != nil {
		return fmt.Errorf("admin service: create: %w")
	}

	return nil
}

func (s *AdminService) GetById(ctx context.Context, id uint) (models.Admin, error) {
	admin, err := s.adminRepo.GetById(ctx, id)
	if err != nil {
		return models.Admin{}, fmt.Errorf("admin service: get by id: %w", err)
	}

	return admin, nil
}
