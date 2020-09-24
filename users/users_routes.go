package users

import (
	"github.com/ElegantSoft/shabahy/common"
	"github.com/ElegantSoft/shabahy/middlewares"
	"github.com/ElegantSoft/shabahy/services"
	"github.com/gin-gonic/gin"
)


var (
	jwtService = services.NewJWTService()
	repository = *NewRepository()
	service    = *NewService(&repository, &jwtService)
	controller = *NewController(&service)
)

func RegisterRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.POST("", controller.CreateUser)
	routerGroup.POST("/auth", controller.Login)
	routerGroup.GET("/me", middlewares.AuthorizeJWT, func(context *gin.Context) {
		id, _ := context.Get(common.K_USER_HEADER)
		context.JSON(200, gin.H{"id": id})
	})
}
