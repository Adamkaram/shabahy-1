package rooms

import (
	"github.com/ElegantSoft/shabahy/common"
	"github.com/ElegantSoft/shabahy/middlewares"
	"github.com/gin-gonic/gin"
)

var (
	crud       = *common.NewCrudRepository("rooms")
	repository = *NewRepository(&crud)
	service    = *NewService(&repository)
	controller = *NewController(&service)
)

func RegisterRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.POST("", controller.create)
	routerGroup.POST("/message/:id", middlewares.AuthorizeJWT, controller.appendMessage)
}
