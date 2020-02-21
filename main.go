package main

import (
	"github.com/gin-gonic/gin"
	"minx.com/demo/app"
)

func main() {
	db := InitDB()
	defer db.Close()

	r := gin.Default()
	app.InitAppModule(db, r)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}