package users

import "github.com/gin-gonic/gin"

var (
	repository = NewRepository()
	service = NewService(repository)
	controller = NewController(service)
)

func RegisterRoutes(routerGroup *gin.RouterGroup)  {
	routerGroup.POST( "/", controller.CreateUser)
	routerGroup.POST( "/auth", controller.Login)
}
