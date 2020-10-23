package users

import (
	"github.com/ElegantSoft/shabahy/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	service Service
}

func (c *Controller) CreateUser(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": common.ValidateErrors(err)})
		return
	}
	token, err := c.service.Create(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"token": token,
		"user": user,
	})
	return

}

func (c *Controller) Login(ctx *gin.Context) {
	var loginData LoginUserDTO
	err := ctx.ShouldBind(&loginData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": common.ValidateErrors(err)})
		return
	}
	token, user, err := c.service.Login(&loginData)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": common.ConvertErrorToString(err, "user not found")})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token, "user": user})

}

func NewController(service Service) *Controller {
	return &Controller{
		service: service,
	}
}
