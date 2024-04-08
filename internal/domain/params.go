package domain

type CarsListParams struct {
	RegNumber string `form:"reg_num"`
	Mark      string `form:"mark"`
	Model     string `form:"model"`
	Year      int    `form:"year"`
	Limit     int    `form:"limit"`
	Offset    int    `form:"offset"`
	OwnerRequest
}

type RegNumberRequest struct {
	RegNumber string `json:"reg_number"`
}

type CarDataUpdatingRequest struct {
	RegNumber string `json:"reg_num"`
	Mark      string `json:"mark"`
	Model     string `json:"model"`
	Year      int    `json:"year"`
	Owner     Owner  `json:"owner"`
}
