package repository

import (
	"book-school/internal/models"
	"context"
	"database/sql"
	"fmt"
)

type AdminRepository struct {
	db *sql.DB
}

func newAdminRepository(db *sql.DB) *AdminRepository {
	return &AdminRepository{
		db: db,
	}
}

func (r *AdminRepository) Create(ctx context.Context, admin models.Admin) error {
	query := fmt.Sprintf("INSERT INTO %s(username, password) VALUES ($1, $2);", adminTable)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("admin repository: create: %w", err)
	}
	defer prep.Close()

	_, err = prep.ExecContext(ctx, admin.Username, admin.Password)
	if err != nil {
		return fmt.Errorf("admin repository: create: %w", err)
	}

	return nil
}

func (r *AdminRepository) GetAdmin(ctx context.Context, name string) (uint, string, error) {
	query := fmt.Sprintf("SELECT id, password FROM %s WHERE username = $1 LIMIT 1", adminTable)
	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, "", fmt.Errorf("admin repo: get admin: %w", err)
	}

	var id uint
	var password string

	if err = prep.QueryRowContext(ctx, name).Scan(&id, &password); err != nil {
		return 0, "", fmt.Errorf("admin repo: get admin: %w", err)
	}

	return id, password, nil
}

func (r *AdminRepository) GetById(ctx context.Context, id uint) (models.Admin, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", adminTable)
	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return models.Admin{}, fmt.Errorf("admin repo: get by id: %w", err)
	}

	var admin models.Admin
	if err = prep.QueryRowContext(ctx, id).Scan(&admin.Id, &admin.Username, &admin.Password); err != nil {
		return models.Admin{}, fmt.Errorf("admin repo: get by id: %w", err)
	}

	return admin, nil
}
