//+build wireinject

package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func inject (db *gorm.DB) AppController{
	wire.Build(ProviderAppRespsitory, ProviderAppService, ProviderAppController)
	return  AppController{}
}

func InitAppModule(db *gorm.DB, r *gin.Engine) {
	controller := inject(db)
	app := r.Group("/app")
	{
		app.GET("/all", controller.FindAll)
	}
}
