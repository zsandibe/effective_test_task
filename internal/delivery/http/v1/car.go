package v1

import (
	"effective/internal/domain"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddCar(c *gin.Context) {
	var inp domain.RegNumberRequest

	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, fmt.Errorf("Incorrect input: %v", err))
		return
	}

	// if err := utils.CheckRegNumber(inp.RegNumber); err != nil {
	// 	newResponse(ctx, http.StatusBadRequest, fmt.Errorf("Incorrect input: %v", err))
	// 	return
	// }

	car := domain.Car{
		RegNum: inp.RegNumber,
	}

	if err := h.service.AddCar(c, car); err != nil {
		fmt.Println("Error tut")
		newResponse(c, http.StatusBadRequest, fmt.Errorf("Failed to add car data: %v", err))
		return
	}

	c.JSON(http.StatusCreated, "Successfully added")
}

func (h *Handler) GetCarById(c *gin.Context) {
	carID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, fmt.Errorf("Incorrect id: %v", err))
		return
	}
	car, err := h.service.GetCarById(c, carID)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, fmt.Errorf("Failed to find car by id: %v", err))
		return
	}

	c.JSON(http.StatusOK, car)
}

func (h *Handler) GetCarsList(c *gin.Context) {
	var input domain.CarsListParams

	if err := c.ShouldBindQuery(&input); err != nil {
		newResponse(c, http.StatusBadRequest, fmt.Errorf("Incorrect input: %v", err))
		return
	}
	fmt.Println(input)
	cars, err := h.service.GetCarsList(c, input)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, fmt.Errorf("Failed to find cars by filter: %v", err))
		return
	}
	if len(cars) == 0 {
		newResponse(c, http.StatusBadRequest, fmt.Errorf("No cars found: %v", errors.New("change params")))
		return
	}
	c.JSON(http.StatusOK, cars)

}

func (h *Handler) UpdateCarInfo(c *gin.Context) {
	var input domain.CarDataUpdatingRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, fmt.Errorf("Incorrect input: %v", err))
		return
	}
	fmt.Println(input)
	carID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, fmt.Errorf("Incorrect id: %v", err))
		return
	}

	if err := h.service.UpdateCarInfo(c, carID, input); err != nil {
		newResponse(c, http.StatusInternalServerError, fmt.Errorf("Failed to update car info: %v", err))
		return
	}
	c.JSON(http.StatusOK, "SUCCESSFULY UPDATED")
}
