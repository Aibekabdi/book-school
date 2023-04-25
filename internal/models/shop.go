package models

type ShopInfo struct {
	Heads []Body `json:"heads"`
	Chest []Body `json:"chest"`
	Legs  []Body `json:"legs"`
	Arms  []Body `json:"arms"`
}

type Body struct {
	Id           uint   `json:"id"`
	Price        uint   `json:"price"`
	Part         string `json:"part"`
	Name         string `json:"name"`
	ImageUrl     string `json:"image_url"`
	ImageIconUrl string `json:"image_icon_url"`
	Buyed        bool   `json:"buyed,omitempty"`
}

const (
	HeadPart  string = "head"
	ChestPart string = "chest"
	LegsPart  string = "legs"
	ArmsPart  string = "arms"
)

type shopCtx string

var (
	CurrentBody shopCtx = "current"
	BodyPart    shopCtx = "body"
	BodyName    shopCtx = "name"
	BodyId      shopCtx = "id"
)
