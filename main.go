package main

import (
	"github.com/ElegantSoft/shabahy/categories"
	"github.com/ElegantSoft/shabahy/db"
	"github.com/ElegantSoft/shabahy/interests"
	"github.com/ElegantSoft/shabahy/users"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	server := gin.Default()
	if err := db.Open(); err != nil {
		log.Fatal(err)
	}

	if err := db.DB.AutoMigrate(
		&users.User{},
		&categories.Category{},
		&interests.Interest{},
	); err != nil {
		log.Fatal(err)
	}

	userGroup := server.Group("users")
	categoriesGroup := server.Group("categories")
	interestsGroup := server.Group("interests")

	users.RegisterRoutes(userGroup)
	categories.RegisterRoutes(categoriesGroup)
	interests.RegisterRoutes(interestsGroup)

	_ = server.Run()

}
