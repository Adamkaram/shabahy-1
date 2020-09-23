package main

import (
	"github.com/ElegantSoft/shabahy/categories"
	"github.com/ElegantSoft/shabahy/common"
	"github.com/ElegantSoft/shabahy/db"
	"github.com/ElegantSoft/shabahy/interests"
	"github.com/ElegantSoft/shabahy/migrations"
	"github.com/ElegantSoft/shabahy/users"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
)

func main() {
	server := gin.Default()
	if err := db.Open(); err != nil {
		log.Fatal(err)
	}
	//custom validators
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("Enum", common.Enum)
	}


	if err := db.DB.AutoMigrate(
		&users.User{},
		&categories.Category{},
		&interests.Interest{},
	); err != nil {
		log.Fatal(err)
	}
	migrations.CreateGenderType()

	userGroup := server.Group("users")
	categoriesGroup := server.Group("categories")
	interestsGroup := server.Group("interests")

	users.RegisterRoutes(userGroup)
	categories.RegisterRoutes(categoriesGroup)
	interests.RegisterRoutes(interestsGroup)
	

	
	_ = server.Run()

}
