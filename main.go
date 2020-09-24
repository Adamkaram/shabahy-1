package main

import (
	"github.com/ElegantSoft/shabahy/categories"
	"github.com/ElegantSoft/shabahy/common"
	"github.com/ElegantSoft/shabahy/db"
	"github.com/ElegantSoft/shabahy/interests"
	"github.com/ElegantSoft/shabahy/users"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}



	server := gin.Default()
	if err := db.Open(); err != nil {
		log.Fatal(err)
	}
	//custom validators
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("Enum", common.Enum)
	}

	//migrations.CreateGenderType()
	//
	//if err := db.DB.AutoMigrate(
	//	&users.User{},
	//	&categories.Category{},
	//	&interests.Interest{},
	//); err != nil {
	//	log.Fatal(err)
	//}

	userGroup := server.Group("users")
	categoriesGroup := server.Group("categories")
	interestsGroup := server.Group("interests")

	users.RegisterRoutes(userGroup)
	categories.RegisterRoutes(categoriesGroup)
	interests.RegisterRoutes(interestsGroup)

	log.Println("main secret", os.Getenv(common.K_JWT_SECRET))


	_ = server.Run()

}
