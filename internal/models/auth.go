package models

type SignInInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Who      string `json:"who"`
}
