//+build wireinject

package app

import (
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"time"
	"google.golang.org/grpc"

	pb "minx.com/demo/grpc"
)

// rest 接口依赖注入
func inject (db *gorm.DB) AppController{
	wire.Build(ProviderAppRepository, ProviderAppService, ProviderAppController)
	return  AppController{}
}

// grpc 接口依赖注入
func inject2(db *gorm.DB) AppGrpc{
	wire.Build(ProviderAppRepository, ProviderAppService, ProviderAppGrpc)
	return AppGrpc{}
}


func InitAppModule(db *gorm.DB, r *gin.Engine, store persistence.CacheStore, s *grpc.Server ) {
	controller := inject(db)
	api := inject2(db)

	// rest api
	app := r.Group("/app")
	{
		app.GET("/all", cache.CachePage(store, time.Minute, controller.FindAll))
	}

	// grpc api
	pb.RegisterAppServiceServer(s, &api)
}
