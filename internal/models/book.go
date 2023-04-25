package models

type Book struct {
	Id        int      `json:"id"`
	Hashed_ID string   `json:"hashed_id"`
	Name      string   `json:"name"`
	Category  string   `json:"category"`
	Class     string   `json:"class"`
	Pages     []string `json:"pages"`
	Preview   string   `json:"preview"`
	Language  string   `json:"language"`
	Audio     []string `json:"audio"`
}

type BooksStruct struct {
	Categories string `json:"categories"`
	Books      []Book `json:"books"`
}
