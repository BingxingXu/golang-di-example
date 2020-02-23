package main

import (
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"

	"minx.com/demo/app"
	_ "minx.com/demo/docs"
)

const port = ":50051"


// @title DI demo
// @version 0.1
// @description 123
// @host http://localhost:8080
func main() {
	db := InitDB()
	defer db.Close()

	r := gin.Default()
	s := grpc.NewServer()
	cache := persistence.NewInMemoryStore(time.Second)

	app.InitAppModule(db, r, cache, s)

	// running grpc server
	go func() {
		log.Printf("starting grpc server at port: %v ", port)

		lis, err := net.Listen("tcp", port)
		if err != nil {
			panic(err)
		}

		if err := s.Serve(lis, ); err != nil {
			panic(err)
		}

	}()

	// running gin server
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	if err := r.Run(); err != nil {
		panic(err)
	}
}