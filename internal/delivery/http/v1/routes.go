package v1

import "github.com/gin-gonic/gin"

func (h *Handler) Routes() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.GET("/swagger/*any")
	api := router.Group("/api/v1")
	{
		car := api.Group("/car")
		{
			car.GET("/:id", h.GetCarById)
			car.POST("/add", h.AddCar)
			car.DELETE("/delete/:id")
			car.GET("/list", h.GetCarsList)
			car.PUT("/update/:id", h.UpdateCarInfo)
		}

	}
	return router
}
