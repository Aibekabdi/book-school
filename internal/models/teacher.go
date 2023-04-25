package models

type Teacher struct {
	Id         uint   `json:"id" swaggerignore:"true"`
	SchoolId   uint   `json:"school_id" swaggerignore:"true"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Private    bool   `json:"private" swaggerignore:"true"`
}

type TeacherUpdate struct {
	TeacherId  uint   `json:"teacher_id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

const TeacherRole = "teacher"

type FullTeacher struct {
	School  School      `json:"school,omitempty" swaggerignore:"true"`
	Teacher Teacher     `json:"teacher"`
	Classes []FullClass `json:"classes"`
}
