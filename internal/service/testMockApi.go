package service

import (
	"effective/internal/domain"
	"fmt"
)

func getCarInfoByRegNum(regNum string) (domain.Car, error) {
	fmt.Println(regNum)
	mockResponse := domain.Car{
		RegNum: "X123XX150",
		Mark:   "Lada",
		Model:  "Vesta",
		Year:   2002,
		Owner: domain.Owner{
			Name:       "Иван",
			Surname:    "Иванов",
			Patronymic: "Иванович",
		},
	}
	return mockResponse, nil
}
