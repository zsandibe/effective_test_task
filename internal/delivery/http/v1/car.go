package v1

import (
	"effective/internal/domain"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddCar godoc
// @Summary Adding new cars
// @Description Adding new cars with his all info
// @Tags car
// @Accept json
// @Produce json
// @Param input body domain.RegNumberRequest true "regNumber"
// @Success 201 {string} string "Successfully added"
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Router /add [post]
func (h *Handler) AddCar(c *gin.Context) {
	var inp domain.RegNumberRequest

	if err := c.BindJSON(&inp); err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf("Incorrect input: %v", err))
		return
	}

	car := domain.Car{
		RegNum: inp.RegNumber,
	}

	if err := h.service.AddCar(c, car); err != nil {
		fmt.Println("Error tut")
		errorResponse(c, http.StatusBadRequest, fmt.Errorf("Failed to add car data: %v", err))
		return
	}

	c.JSON(http.StatusCreated, "Successfully added")
}

// GetCarById godoc
// @Summary Get car info by id
// @Description Getting car info by id
// @Tags car
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} domain.Car
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Router /{id} [get]
func (h *Handler) GetCarById(c *gin.Context) {
	carID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf("Incorrect id: %v", err))
		return
	}
	car, err := h.service.GetCarById(c, carID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, fmt.Errorf("Failed to find car by id: %v", err))
		return
	}

	c.JSON(http.StatusOK, car)
}

// GetCarList godoc
// @Summary Get cars list by filter
// @Description Getting cars info by filter
// @Tags car
// @Accept json
// @Produce json
// @Param reg_num query string false "Car reg number"
// @Param mark query string false "Car mark"
// @Param model query string false "Car model"
// @Param limit query int false "Limit for pagination"
// @Param offset query string false "Offset for pagination"
// @Param name query string false "Owner`s name"
// @Param surname query string false "Owner`s surname"
// @Param patronymic query string false "Owner`s patronymic"
// @Success 200 {object} []domain.Car
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Router /list [get]
func (h *Handler) GetCarsList(c *gin.Context) {
	var input domain.CarsListParams

	if err := c.ShouldBindQuery(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf("Incorrect input: %v", err))
		return
	}
	fmt.Println(input)
	cars, err := h.service.GetCarsList(c, input)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, fmt.Errorf("Failed to find cars by filter: %v", err))
		return
	}
	if len(cars) == 0 {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf("No cars found: %v", errors.New("change params")))
		return
	}
	c.JSON(http.StatusOK, cars)

}

// UpdateCarInfo godoc
// @Summary Update car information
// @Description Updating car details by ID
// @Tags car
// @Accept json
// @Produce json
// @Param id path int true "Car ID"
// @Param   car body domain.CarDataUpdatingRequest true "Update Car Request"
// @Success 200 {string} string "Succesfully updated"
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Router /update/{id} [put]
func (h *Handler) UpdateCarInfo(c *gin.Context) {
	var input domain.CarDataUpdatingRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf("Incorrect input: %v", err))
		return
	}
	fmt.Println(input)
	carID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf("Incorrect id: %v", err))
		return
	}

	if err := h.service.UpdateCarInfo(c, carID, input); err != nil {
		errorResponse(c, http.StatusInternalServerError, fmt.Errorf("Failed to update car info: %v", err))
		return
	}
	c.JSON(http.StatusOK, "Successfully updated")
}

// DeleteCarById godoc
// @Summary Delete a car
// @Description Delete a car by Id
// @Tags car
// @Accept  json
// @Produce  json
// @Param   id path int true "Car ID"
// @Success 200 {string} string "Successfully deleted"
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /cars/{id} [delete]
func (h *Handler) DeleteCarById(c *gin.Context) {
	carID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf("Incorrect id: %v", err))
		return
	}
	if err := h.service.DeleteCarById(c, carID); err != nil {
		errorResponse(c, http.StatusInternalServerError, fmt.Errorf("Failed to find car by id: %v", err))
		return
	}

	c.JSON(http.StatusOK, "Successfully deleted")
}
