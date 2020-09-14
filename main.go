package main

import (
	"github.com/ElegantSoft/shabahy/db"
	"github.com/ElegantSoft/shabahy/users"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	server := gin.Default()
	if err := db.Open(); err != nil {
		log.Fatal(err)
	}

	db.DB.AutoMigrate(&users.User{})

	userGroup := server.Group("users")
	users.RegisterRoutes(userGroup)

	server.Run()


}
