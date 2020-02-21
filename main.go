package main

import (
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"

	"minx.com/demo/app"
)

const port = ":50051"

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

	// running main server
	if err := r.Run(); err != nil {
		panic(err)
	}
}