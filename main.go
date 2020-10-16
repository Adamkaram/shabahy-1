package main

import (
	"fmt"
	"github.com/ElegantSoft/shabahy/categories"
	"github.com/ElegantSoft/shabahy/common"
	"github.com/ElegantSoft/shabahy/db"
	"github.com/ElegantSoft/shabahy/interests"
	"github.com/ElegantSoft/shabahy/messages"
	"github.com/ElegantSoft/shabahy/migrations"
	"github.com/ElegantSoft/shabahy/rooms"
	"github.com/ElegantSoft/shabahy/users"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	socketio "github.com/googollee/go-socket.io"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	socketServer, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal("error in socket.io", err)
	}
	socketServer.OnConnect("/", func(s socketio.Conn) error {
		
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	socketServer.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", msg)
	})


	socketServer.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})

	socketServer.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	socketServer.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	socketServer.OnDisconnect("", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	go socketServer.Serve()
	defer socketServer.Close()



	server := gin.Default()
	if err := db.Open(); err != nil {
		log.Fatal(err)
	}
	//custom validators
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("Enum", common.Enum)
	}

	migrations.CreateGenderType()

	if err := db.DB.AutoMigrate(
		&users.User{},
		&categories.Category{},
		&interests.Interest{},
		&rooms.Room{},
		&messages.Message{},
	); err != nil {
		log.Fatal(err)
	}

	userGroup := server.Group("users")
	categoriesGroup := server.Group("categories")
	interestsGroup := server.Group("interests")
	roomsGroup := server.Group("rooms")
	messagesGroup := server.Group("messages")


	server.GET("/socket.io/*any", gin.WrapH(socketServer))
	server.POST("/socket.io/*any", gin.WrapH(socketServer))

	users.RegisterRoutes(userGroup)
	categories.RegisterRoutes(categoriesGroup)
	interests.RegisterRoutes(interestsGroup)
	rooms.RegisterRoutes(roomsGroup)
	messages.RegisterRoutes(messagesGroup)
	server.StaticFS("/public", http.Dir("./assets"))

	log.Println("main secret", os.Getenv(common.KJwtSecret))


	_ = server.Run()

}
