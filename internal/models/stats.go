package models

type Stats struct {
	StudentFirstName   string `json:"student_first_name"`
	StudentSecondName  string `json:"student_second_name"`
	Grade              string `json:"grade"`
	Name               string `json:"name"`
	StudentId          uint   `json:"student_id"`
	BookPoints         uint   `json:"book_points"`
	AudioPoints        uint   `json:"audio_points"`
	TestPoints         uint   `json:"test_points"`
	CreativeTaskPoints uint   `json:"creative_task_points"`
	OpenPoints      uint   `json:"open_points"`
	TotalPoints        uint   `json:"total_points"`
}
