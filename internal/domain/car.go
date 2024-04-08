package domain

type Car struct {
	Id     int    `json:"id"`
	RegNum string `json:"reg_num"`
	Mark   string `json:"mark"`
	Model  string `json:"model"`
	Year   int    `json:"year"`
	Owner  Owner  `json:"owner"`
}
