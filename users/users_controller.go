package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	service Service
}

func (c Controller) CreateUser(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBind(&user); err == nil {
		_, created := c.service.Create(&user)
		ctx.JSON(http.StatusCreated, created)
	} else {
		ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
}

func NewController(service *Service) *Controller {
	return &Controller{
		service: *service,
	}
}
