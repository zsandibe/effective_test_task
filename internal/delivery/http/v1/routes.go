package v1

import (
	_ "effective/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     //
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func (h *Handler) Routes() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api/v1")
	{

		api.GET("/:id", h.GetCarById)
		api.POST("/add", h.AddCar)
		api.DELETE("/delete/:id", h.DeleteCarById)
		api.GET("/list", h.GetCarsList)
		api.PUT("/update/:id", h.UpdateCarInfo)

	}
	return router
}
