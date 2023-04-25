package models

type Student struct {
	Id         uint   `json:"id" swaggerignore:"true"`
	ClassId    uint   `json:"class_id"`
	Points     uint   `json:"points"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	CountPass  uint   `json:"count_pass"`
}

type StudentUpdate struct {
	StudentId  uint   `json:"student_id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

const StudentRole = "student"
