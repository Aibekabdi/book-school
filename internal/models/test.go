package models

type Test struct {
	Id        uint       `json:"id" swaggerignore:"true"`
	BookId    uint       `json:"book_id"`
	Lang      string     `json:"lang"`
	Questions []Question `json:"questions"`
}

type TestForTeacher struct {
	Student Student `json:"student"`
	Test    Test    `json:"test"`
}

type CompleteTest struct {
	Id        uint              `json:"id" swaggerignore:"true"`
	StudentId uint              `json:"student_id" swaggerignore:"true"`
	TestId    uint              `json:"test_id"`
	Points    uint              `json:"points"`
	Answers   []CompleteAnswers `json:"answers"`
}

type CompleteAnswers struct {
	Id             uint `json:"id" swaggerignore:"true"`
	CompleteTestId uint `json:"complete_test_id" swaggerignore:"true"`
	QuestionId     uint `json:"question_id"`
	AnswerId       uint `json:"answer_id"`
}

type CompleteTestResp struct {
	TestId              int    `json:"test_id"`
	CorrectAnswers      []bool `json:"correct_answers"`
	Points              int    `json:"points"`
	CorrectAnswersCount int    `json:"correct_answers_count"`
}
