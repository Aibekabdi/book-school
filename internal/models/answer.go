package models

import "mime/multipart"

type Answer struct {
	Id               uint   `json:"id" swaggerignore:"true"`
	QuestionId       uint   `json:"question_id" swaggerignore:"true"`
	Image            string `json:"image" swaggerignore:"true"`
	Answer           string `json:"answer"`
	Audio            string `json:"audio"`
	WithImage        bool   `json:"with_image"`
	Correct          bool   `json:"correct"`
	IsStudentsAnswer bool   `json:"is_students_answer"`
}

type PassCreativeTask struct {
	Id         int    `json:"id"`
	StudentId  uint   `json:"student_id"`
	BookId     int    `json:"book_id"`
	QuestionId int    `json:"question_id"`
	Question   string `json:"question"`
	Answer     string `json:"Answer"`
	Img        multipart.File
	IsArt      bool `json:"is_art"`
}

type CreativeTaskData struct {
	QuestionId int    `json:"question_id"`
	StudentId  uint   `json:"student_id"`
	Username   string `json:"username"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
}
