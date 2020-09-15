package categories

import (
	"github.com/gin-gonic/gin"
)

var (
	repository = *NewRepository()
	service = *NewService(&repository)
	controller = *NewController(&service)
)

func RegisterRoutes(routerGroup *gin.RouterGroup)  {
	routerGroup.GET( "/:id", controller.Find)
	routerGroup.POST( "/", controller.Create)
	routerGroup.DELETE( "/:id", controller.Delete)
	routerGroup.PUT( "/:id", controller.Update)
}
