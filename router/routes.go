package router

import (
	"github.com/gin-gonic/gin"
	"github.com/luizfernandoOliveiraa/Go-Api/handler"
)

func initializeRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.GET("/oppening", handler.ShowOpeningHanlder())

		v1.POST("/oppening", handler.CreateOpeningHanlder())

		v1.DELETE("/oppening", handler.DeleteOpeningHanlder())

		v1.PUT("/oppening", handler.UpdateOpeningHanlder())

		v1.GET("/oppenings", handler.ListOpeningsHanlder())

	}

}
