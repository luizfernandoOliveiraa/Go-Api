package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizfernandoOliveiraa/Go-Api/schemas"
)

func ListOpeningsHanlder(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSucess(ctx, "lists-opening", openings)
}
