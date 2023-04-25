package models

type School struct {
	Id         uint   `json:"id" swaggerignore:"true"`
	ClassCount uint   `json:"class_count"`
	Name       string `json:"name"`
	Password   string `json:"password"`
}

type SchoolUpdate struct {
	SchoolId   uint   `json:"school_id"`
	ClassCount uint   `json:"class_count"`
	Name       string `json:"name"`
	Password   string `json:"password"`
}

const SchoolRole = "school"

type FullSchool struct {
	School        School        `json:"school"`
	TotalClasses  int           `json:"total_classes"`
	TotalStudents int           `json:"total_students"`
	Teachers      []FullTeacher `json:"teachers"`
}
