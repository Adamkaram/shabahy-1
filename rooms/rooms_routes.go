package rooms

import (
	"github.com/ElegantSoft/shabahy/common"
	"github.com/ElegantSoft/shabahy/middlewares"
	"github.com/gin-gonic/gin"
)

var (
	crud       = *common.NewCrudRepository("rooms")
	repository = *NewRepository(&crud)
	RoomService    = *NewService(&repository)
	controller = *NewController(&RoomService)
)

func RegisterRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.POST("", controller.create)
	routerGroup.POST("/message/:hash", middlewares.AuthorizeJWT, controller.appendMessage)
	routerGroup.GET("/get/:id", middlewares.AuthorizeJWT, controller.find)
	//routerGroup.GET("/test", func(context *gin.Context) {
	//	var result []Room
	//	db.DB.Model(&Room{}).Preload(RoomSchema.Users, func(db2 *gorm.DB) *gorm.DB {
	//		return db2.Select("name", "phone")
	//	}).Select("hash", "id").Find(&result)
	//	context.JSON(200, gin.H{"data": result})
	//})
}
