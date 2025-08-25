package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/luizfernandoOliveiraa/Go-Api/docs"
	"github.com/luizfernandoOliveiraa/Go-Api/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {

	//Initialize Handler
	handler.InitalizeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		v1.GET("/oppening", handler.ShowOpeningHanlder)
		v1.POST("/oppening", handler.CreateOpeningHanlder)
		v1.DELETE("/oppening", handler.DeleteOpeningHanlder)
		v1.PUT("/oppening", handler.UpdateOpeningHanlder)
		v1.GET("/oppenings", handler.ListOpeningsHanlder)

	}

	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
