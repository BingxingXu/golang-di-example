// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Injectors from app.module.go:

func inject(db *gorm.DB) AppController {
	appQuerySet := ProviderAppRespsitory(db)
	appService := ProviderAppService(appQuerySet)
	appController := ProviderAppController(appService)
	return appController
}

// app.module.go:

func InitAppModule(db *gorm.DB, r *gin.Engine) {
	controller := inject(db)
	app := r.Group("/app")
	{
		app.GET("/all", controller.FindAll)
	}
}