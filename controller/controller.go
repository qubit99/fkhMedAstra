package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
func (c *Controller) UpdateUserProfile(ctx *gin.Context) {
	updateProfileReq := models.UserProfile{}
	err := ctx.ShouldBindJSON(&updateProfileReq)
	updateProfileReq.Username = ctx.Param("username")
	if len(updateProfileReq.Username) == 0 {
		ctx.JSON(http.StatusBadRequest, "No username for update")
		return
	}
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	err = c.svc.UpdateUserProfile(&updateProfileReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, "Updated")
}
func (c *Controller) FindUserByUsername(ctx *gin.Context) {
	username := ctx.Param("username")
	if len(username) == 0 {
		ctx.JSON(http.StatusBadRequest, "No username supplied")
	}
	user, err := c.svc.GetUserProfile(username)
	if err == gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusNotFound, err)
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, user)

}
