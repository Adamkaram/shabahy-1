package main

import (
	"github.com/ElegantSoft/shabahy/categories"
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

	_ = db.DB.AutoMigrate(&users.User{}, &categories.Category{})

	userGroup := server.Group("users")
	categoriesGroup := server.Group("categories")

	users.RegisterRoutes(userGroup)
	categories.RegisterRoutes(categoriesGroup)

	server.Run()


}
