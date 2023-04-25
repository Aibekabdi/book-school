package models

type Admin struct {
	Id       uint   `json:"id" swaggerignore:"true"`
	Username string `json:"username"`
	Password string `json:"password"`
}

const AdminRole = "admin"
