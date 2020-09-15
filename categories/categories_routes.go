package categories

import (
	"github.com/ElegantSoft/shabahy/common"
	"github.com/gin-gonic/gin"
)

var (
	crud = common.CrudRepository{Model: Category{}}
	repository = *NewRepository(&crud)
	service = *NewService(&repository)
	controller = *NewController(&service)
)

func RegisterRoutes(routerGroup *gin.RouterGroup)  {
	routerGroup.GET( "/:id", controller.Find)
	routerGroup.POST( "/", controller.Create)
	routerGroup.DELETE( "/:id", controller.Delete)
	routerGroup.PUT( "/:id", controller.Update)
}
