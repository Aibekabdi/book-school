package models

type Question struct {
	Id        uint     `json:"id" swaggerignore:"true"`
	TestId    uint     `json:"test_id" swaggerignore:"true"`
	WithImage bool     `json:"with_image"`
	Image     string   `json:"image"`
	Question  string   `json:"question" swaggerignore:"true"`
	Audio     string   `json:"audio"`
	Answers   []Answer `json:"answers"`
}

type CreativeTask struct {
	Id       uint   `json:"id"`
	Category string `json:"category"`
	Question string `json:"question"`
	IsArt    bool   `json:"is_art"`
	Audio    string `json:"audio"`
}
