package rooms

import (
	"github.com/ElegantSoft/shabahy/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	service *Service
}

func (s *Controller) create(ctx *gin.Context) {
	var users struct{ Users []uint `json:"users" binding:"required"` }
	if err := ctx.ShouldBind(&users); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err, found := s.service.create(users.Users)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": &found})
}

func (s *Controller) appendMessage(ctx *gin.Context) {
	var message Message
	var params struct{ID uint `uri:"id" binding:"required"`}
	idHeader, _ := ctx.Get(common.KUserHeader)
	userId := common.GetIdFromCtx(idHeader)
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": common.ValidateErrors(err)})
		return
	}
	if err := ctx.ShouldBind(&message); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": common.ValidateErrors(err)})
		return
	}
	err := s.service.appendMessage(params.ID, &message, userId)
	if err != nil {
	    ctx.JSON(http.StatusBadRequest, gin.H{"error": common.ValidateErrors(err)})
	    return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "created", "data": message})

}

//func (s *Controller) find(ctx *gin.Context) {
//	var item ById
//	if err := ctx.ShouldBindUri(&item); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	err, found := s.service.find(item.ID)
//	if err != nil {
//		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
//		return
//	}
//	ctx.JSON(http.StatusOK, gin.H{"data": &found})
//}

//func (s *Controller) paginate(ctx *gin.Context) {
//	err, found := s.service.paginate()
//	if err != nil {
//		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
//		return
//	}
//	ctx.JSON(http.StatusOK, gin.H{"data": &found})
//}


//func (s *Controller) Delete(ctx *gin.Context) {
//	var item ById
//	if err := ctx.ShouldBindUri(&item); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	err := s.service.delete(item.ID)
//	if err != nil {
//		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
//		return
//	}
//	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
//}

//func (s *Controller) Update(ctx *gin.Context) {
//	var item Room
//	var byId ById
//	if err := ctx.ShouldBind(&item); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	if err := ctx.ShouldBindUri(&byId); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	err := s.service.update(&item, byId.ID)
//	if err != nil {
//		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
//		return
//	}
//	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
//}

func NewController(service *Service) *Controller {
	return &Controller{
		service: service,
	}
}
