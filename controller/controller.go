package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"medastra/models"
	"medastra/service"
	"net/http"
)

type Controller struct {
	svc service.Service
}

func NewController(svc service.Service) *Controller {
	return &Controller{svc: svc}
}
func (c *Controller) CreateUserProfile(ctx *gin.Context) {
	createProfileReq := models.UserProfile{}
	err := ctx.ShouldBindJSON(&createProfileReq)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	err = c.svc.CreateUserProfile(&createProfileReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusCreated, createProfileReq)
}
