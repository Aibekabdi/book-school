package repository

import (
	"book-school/internal/models"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type ShopRepository struct {
	db *sql.DB
}

func newShopRepository(db *sql.DB) *ShopRepository {
	return &ShopRepository{
		db: db,
	}
}

func (r *ShopRepository) Buy(ctx context.Context, studentId, bodyId uint) error {
	query := fmt.Sprintf("INSERT INTO %s(student_id, body_id) VALUES ($1, $2);", buyedTable)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("shop repository: buy: %w", err)
	}
	defer prep.Close()

	_, err = prep.ExecContext(ctx, studentId, bodyId)
	if err != nil {
		return fmt.Errorf("shop repository: buy: %w", err)
	}

	return nil
}

func (r *ShopRepository) GetOneBy(ctx context.Context) (models.Body, error) {
	args := []interface{}{}
	argsStr := []string{}
	argsNum := 1

	id := ctx.Value(models.BodyId)
	if id != nil {
		args = append(args, id.(uint))
		argsStr = append(argsStr, fmt.Sprintf("id = $%d", argsNum))
		argsNum++
	}

	name := ctx.Value(models.BodyName)
	if name != nil {
		args = append(args, name.(string))
		argsStr = append(argsStr, fmt.Sprintf("name = $%d", argsNum))
		argsNum++
	}

	whereCondition := ""

	if len(argsStr) != 0 {
		whereCondition = " WHERE " + strings.Join(argsStr, " AND ")
	}

	query := fmt.Sprintf("SELECT * FROM %s%s;", bodyTable, whereCondition)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return models.Body{}, fmt.Errorf("shop repository: get all: %w", err)
	}
	defer prep.Close()

	var body models.Body

	err = prep.QueryRowContext(ctx, args...).Scan(&body.Id, &body.Part, &body.Name, &body.ImageUrl, &body.ImageIconUrl, &body.Price)
	if err != nil {
		return models.Body{}, fmt.Errorf("shop repository: get all: %w", err)
	}

	return body, nil
}

func (r *ShopRepository) Create(ctx context.Context, body models.Body) (uint, error) {
	query := fmt.Sprintf("INSERT INTO %s(part, name, img_url, img_icon_url, price) VALUES ($1, $2, $3, $4, $5) RETURNING id;", bodyTable)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, fmt.Errorf("shop repository: create: %w", err)
	}
	defer prep.Close()

	var id uint

	err = prep.QueryRowContext(ctx, body.Part, body.Name, body.ImageUrl, body.ImageIconUrl, body.Price).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("shop repository: create: %w", err)
	}

	return id, nil
}

func (r *ShopRepository) GetAll(ctx context.Context) ([]models.Body, error) {
	args := []interface{}{}
	argsNum := 1

	whereCondition := ""

	bodyPart := ctx.Value(models.BodyPart)
	if bodyPart != nil {
		whereCondition = fmt.Sprintf(" WHERE part=$%d", argsNum)
		args = append(args, bodyPart.(string))
		argsNum++
	}

	studentId := ctx.Value(models.StudentId)
	if studentId != nil {
		whereCondition += fmt.Sprintf(" AND id IN (SELECT body_id FROM %s WHERE student_id = $%d)", buyedTable, argsNum)
		args = append(args, studentId.(uint))
		argsNum++
	}

	currentBody := ctx.Value(models.CurrentBody)
	if currentBody != nil {
		whereCondition = strings.ReplaceAll(whereCondition, buyedTable, currentBodyTable)
		whereCondition = strings.ReplaceAll(whereCondition, "body_id", bodyPart.(string)+"_id")
	}

	query := fmt.Sprintf("SELECT * FROM %s%s;", bodyTable, whereCondition)

	prep, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("shop repository: get all: %w", err)
	}
	defer prep.Close()

	rows, err := prep.QueryContext(ctx, args...)
	if err != nil {
		return nil, fmt.Errorf("shop repository: get all: %w", err)
	}
	defer rows.Close()

	var bodies []models.Body

	for rows.Next() {
		var body models.Body

		err = rows.Scan(&body.Id, &body.Part, &body.Name, &body.ImageUrl, &body.ImageIconUrl, &body.Price)
		if err != nil {
			return nil, fmt.Errorf("shop repository: get all: %w", err)
		}

		bodies = append(bodies, body)
	}

	return bodies, nil
}
