package models

type CheckCreativePass struct {
	TeacherId uint   `json:"teacher_id"`
	StudentId uint   `json:"student_id"`
	AnswerId  uint   `json:"answer_id"`
	Comment   string `json:"comment"`
	Point     uint   `json:"point"`
	Audio     string `json:"audio"`
}

type CreativeNotifications struct {
	Answer   string `json:"answer"`
	Comment  string `json:"comment"`
	Audio    string `json:"audio"`
	Question string `json:"question"`
	BookName string `json:"book_name"`
}
