package models

type Class struct {
	Id        uint   `json:"id" swaggerignore:"true"`
	SchoolId  uint   `json:"school_id" swaggerignore:"true"`
	TeacherId uint   `json:"teacher_id"`
	Grade     string `json:"grade"`
	Name      string `json:"name"`
}

type FullClass struct {
	Class    Class     `json:"class"`
	Students []Student `json:"students"`
}

type ClassStats struct {
	Class Class   `json:"class"`
	Stats []Stats `json:"stats"`
}

type classCtx string

var ClassGrade classCtx = "class_grade"

type categoriesCtx string

var Categoryies categoriesCtx = "categories"
