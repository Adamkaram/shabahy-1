package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	service *Service
}

func (c *Controller) CreateUser(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBind(&user); err == nil {
		_, created := c.service.Create(&user)
		ctx.JSON(http.StatusCreated, created)
	} else {
		ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
}

func (c *Controller) Login(ctx *gin.Context) {
	var loginData LoginUserDTO
	err := ctx.ShouldBind(&loginData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		loginError, token := c.service.Login(&loginData)
		if loginError != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": loginError.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"token": token})
		}
	}
}

func NewController(service *Service) *Controller {
	return &Controller{
		service: service,
	}
}
