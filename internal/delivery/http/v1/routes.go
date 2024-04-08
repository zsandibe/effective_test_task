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
			car.GET("/:id")
			car.POST("/add")
			car.DELETE("/delete/:id")
			car.GET("/list")
			car.PUT("/update/:id")
		}

	}
	return router
}
