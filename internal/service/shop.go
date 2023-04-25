package service

import (
	"book-school/internal/models"
	"book-school/internal/repository"
	"book-school/pkg/config"
	"context"
	"fmt"
	"mime/multipart"
	"strconv"
	"strings"
)

type ShopService struct {
	shopRepo    repository.Shop
	studentRepo repository.Student

	cfg *config.Conf
}

func newShopService(shopRepo repository.Shop, studentRepo repository.Student, cfg *config.Conf) *ShopService {
	return &ShopService{
		shopRepo:    shopRepo,
		studentRepo: studentRepo,
		cfg:         cfg,
	}
}

func (s *ShopService) validate(body models.Body) error {
	if len(body.Name) < 3 || len(body.Name) > 50 || strings.ReplaceAll(body.Name, " ", "") == "" {
		return fmt.Errorf("validate: invalid body name")
	}

	if body.Part != models.ArmsPart && body.Part != models.LegsPart && body.Part != models.ChestPart && body.Part != models.HeadPart {
		return fmt.Errorf("validate: invalid body part")
	}

	if body.Price <= 0 {
		return fmt.Errorf("validate: invalid body price")
	}

	return nil
}

func (s *ShopService) UpdateCurrentBody(ctx context.Context, studentId, fromId, toId uint) error {
	fromBody, err := s.shopRepo.GetOneBy(context.WithValue(ctx, models.BodyId, fromId))
	if err != nil {
		return fmt.Errorf("shop service: update current body: %w", err)
	}

	toBody, err := s.shopRepo.GetOneBy(context.WithValue(ctx, models.BodyId, toId))
	if err != nil {
		return fmt.Errorf("shop service: update current body: %w", err)
	}

	if fromBody.Part != toBody.Part {
		return fmt.Errorf("shop service: update current body: cant chage different parts")
	}

	info, err := s.GetAllBuyed(ctx, studentId)
	if err != nil {
		return fmt.Errorf("shop service: update current body: %w", err)
	}

	temp := []models.Body{}

	temp = append(temp, info.Heads...)
	temp = append(temp, info.Legs...)
	temp = append(temp, info.Arms...)
	temp = append(temp, info.Chest...)

	if !s.isIn(temp, toBody) {
		return fmt.Errorf("shop service: update: you dont own this part")
	}

	err = s.studentRepo.UpdateBody(ctx, studentId, fromId, toId, fromBody.Part)
	if err != nil {
		return fmt.Errorf("shop service: update current body: %w", err)
	}

	return nil
}

func (s *ShopService) Create(ctx context.Context, image, imageIcon *multipart.FileHeader, price, name, part string) (models.Body, error) {
	var body models.Body

	priceInt, err := strconv.Atoi(price)
	if err != nil {
		return models.Body{}, fmt.Errorf("shop service: create: %w", err)
	}

	body.Name = name
	body.Price = uint(priceInt)
	body.Part = part

	err = s.validate(body)
	if err != nil {
		return models.Body{}, fmt.Errorf("shop service: create: %w", err)
	}

	url, err := saveFiles("./static/shop/view/", body.Part, image)
	if err != nil {
		return models.Body{}, fmt.Errorf("shop service: create: %w", err)
	}

	body.ImageUrl = "http://" + s.cfg.Api.Host + ":" + s.cfg.Api.Port + url[1:]

	iconUrl, err := saveFiles("./static/shop/icon/", body.Part, imageIcon)
	if err != nil {
		return models.Body{}, fmt.Errorf("shop service: create: %w", err)
	}

	body.ImageUrl = "http://" + s.cfg.Api.Host + ":" + s.cfg.Api.Port + iconUrl[1:]

	body.Id, err = s.shopRepo.Create(ctx, body)
	if err != nil {
		return models.Body{}, fmt.Errorf("shop service: create: %w", err)
	}

	return body, nil
}

func (s *ShopService) GetAllBuyed(ctx context.Context, studentId uint) (models.ShopInfo, error) {
	var (
		info models.ShopInfo
		err  error
	)

	ctx = context.WithValue(ctx, models.StudentId, studentId)

	info.Heads, err = s.shopRepo.GetAll(context.WithValue(ctx, models.BodyPart, models.HeadPart))
	if err != nil {
		return models.ShopInfo{}, fmt.Errorf("shop service: get all buyed: %w", err)
	}

	info.Arms, err = s.shopRepo.GetAll(context.WithValue(ctx, models.BodyPart, models.ArmsPart))
	if err != nil {
		return models.ShopInfo{}, fmt.Errorf("shop service: get all buyed: %w", err)
	}

	info.Chest, err = s.shopRepo.GetAll(context.WithValue(ctx, models.BodyPart, models.ChestPart))
	if err != nil {
		return models.ShopInfo{}, fmt.Errorf("shop service: get all buyed: %w", err)
	}

	info.Legs, err = s.shopRepo.GetAll(context.WithValue(ctx, models.BodyPart, models.LegsPart))
	if err != nil {
		return models.ShopInfo{}, fmt.Errorf("shop service: get all buyed: %w", err)
	}

	return info, nil
}

func (s *ShopService) GetCurrentBody(ctx context.Context, studentId uint) (models.ShopInfo, error) {
	info, err := s.GetAllBuyed(context.WithValue(ctx, models.CurrentBody, true), studentId)
	if err != nil {
		return models.ShopInfo{}, fmt.Errorf("shop service: get current body: %w", err)
	}

	return info, err
}

func (s *ShopService) Buy(ctx context.Context, studentId uint, bodyId uint) (uint, error) {
	body, err := s.shopRepo.GetOneBy(context.WithValue(ctx, models.BodyId, bodyId))
	if err != nil {
		return 0, fmt.Errorf("shop service: buy: %w", err)
	}

	student, err := s.studentRepo.GetById(ctx, studentId)
	if err != nil {
		return 0, fmt.Errorf("shop service: buy: %w", err)
	}

	if student.Points < body.Price {
		return 0, fmt.Errorf("shop service: buy: not enough coins")
	}

	info, err := s.GetAllBuyed(ctx, student.Id)
	if err != nil {
		return 0, fmt.Errorf("shop service: buy: %w", err)
	}

	temp := []models.Body{}

	temp = append(temp, info.Heads...)
	temp = append(temp, info.Legs...)
	temp = append(temp, info.Arms...)
	temp = append(temp, info.Chest...)

	if s.isIn(temp, body) {
		return 0, fmt.Errorf("shop service: buy: you already bought this part")
	}

	err = s.studentRepo.TakePoints(ctx, student.Id, body.Price)
	if err != nil {
		return 0, fmt.Errorf("shop service: buy: %w", err)
	}

	err = s.shopRepo.Buy(ctx, student.Id, body.Id)
	if err != nil {
		err = s.studentRepo.GivePoints(ctx, student.Id, body.Price)
		if err != nil {
			return 0, fmt.Errorf("shop service: buy: %w", err)
		}
		return 0, fmt.Errorf("shop service: buy: %w", err)
	}

	return student.Points - body.Price, nil
}

func (s *ShopService) isIn(arr []models.Body, body models.Body) bool {
	for _, v := range arr {
		if v.Id == body.Id {
			return true
		}
	}
	return false
}

func (s *ShopService) GetAll(ctx context.Context, studentId uint) (models.ShopInfo, error) {
	var (
		info models.ShopInfo
		err  error
	)

	buyed, err := s.GetAllBuyed(ctx, studentId)
	if err != nil {
		return models.ShopInfo{}, fmt.Errorf("shop service: get all: %w", err)
	}

	info.Heads, err = s.shopRepo.GetAll(context.WithValue(ctx, models.BodyPart, models.HeadPart))
	if err != nil {
		return models.ShopInfo{}, fmt.Errorf("shop service: get all: %w", err)
	}

	for i, head := range info.Heads {
		for _, buyedHead := range buyed.Heads {
			if head.Id == buyedHead.Id {
				info.Heads[i].Buyed = true
				break
			}
		}
	}

	info.Arms, err = s.shopRepo.GetAll(context.WithValue(ctx, models.BodyPart, models.ArmsPart))
	if err != nil {
		return models.ShopInfo{}, fmt.Errorf("shop service: get all: %w", err)
	}

	for i, arms := range info.Arms {
		for _, buyedArms := range buyed.Arms {
			if arms.Id == buyedArms.Id {
				info.Arms[i].Buyed = true
				break
			}
		}
	}

	info.Chest, err = s.shopRepo.GetAll(context.WithValue(ctx, models.BodyPart, models.ChestPart))
	if err != nil {
		return models.ShopInfo{}, fmt.Errorf("shop service: get all: %w", err)
	}

	for i, chest := range info.Chest {
		for _, buyedChest := range buyed.Chest {
			if chest.Id == buyedChest.Id {
				info.Chest[i].Buyed = true
				break
			}
		}
	}

	info.Legs, err = s.shopRepo.GetAll(context.WithValue(ctx, models.BodyPart, models.LegsPart))
	if err != nil {
		return models.ShopInfo{}, fmt.Errorf("shop service: get all: %w", err)
	}

	for i, legs := range info.Legs {
		for _, buyedLegs := range buyed.Legs {
			if legs.Id == buyedLegs.Id {
				info.Legs[i].Buyed = true
				break
			}
		}
	}

	return info, nil
}
