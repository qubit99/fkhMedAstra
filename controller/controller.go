package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"medastra/models"
	"medastra/service"
	"net/http"
	"strconv"
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

func (c *Controller) Login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	err = c.svc.Login(user.Username, user.Password)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}
	ctx.JSON(http.StatusOK, "Login successful")
}

func (c *Controller) CreateAccount(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	err = c.svc.CreateAccount(&user)
	ctx.JSON(http.StatusCreated, user)
}

func (c *Controller) CreateBooking(ctx *gin.Context) {
	var booking models.Booking
	err := ctx.ShouldBindJSON(&booking)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusAccepted, err)
		return
	}
	err = c.svc.CreateBooking(&booking)
	ctx.JSON(http.StatusCreated, booking)
}

func (c *Controller) GetBookings(ctx *gin.Context) {
	username := ctx.Param("username")
	if len(username) == 0 {
		ctx.JSON(http.StatusBadRequest, "No username supplied")
	}
	bookings, err := c.svc.GetBookings(username)
	if err == gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusNotFound, err)
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, bookings)

func (c *Controller) GetDoctors(ctx *gin.Context) {
	var searchReq models.DoctorSearchRequest
	err := ctx.ShouldBindJSON(&searchReq)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	doctorsResponse, err := c.svc.GetDoctors(&searchReq)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, doctorsResponse)
}

func (c *Controller) BookSlot(ctx *gin.Context) {
	username := ctx.Param("username")
	slotIdStr := ctx.Param("slotId")

	slotId, err := strconv.Atoi(slotIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	err = c.svc.BookSlot(username, slotId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, "Booked")
}

func (c *Controller) UserBookings(ctx *gin.Context) {
	username := ctx.Param("username")
	bookings, err := c.svc.UserBookings(username)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, bookings)
}

func (c *Controller) DoctorBookings(ctx *gin.Context) {
	doctoridStr := ctx.Param("id")
	doctorId, err := strconv.Atoi(doctoridStr)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	bookings, err := c.svc.DoctorBookings(doctorId)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, bookings)
}
