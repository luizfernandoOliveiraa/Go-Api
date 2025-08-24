package handler

import (
	"github.com/luizfernandoOliveiraa/Go-Api/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitalizeHandler(){
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}