package routes

import (
	"github.com/gin-gonic/gin"
	"medastra/controller"
	"net/http"
)

func NewRouter(controller *controller.Controller) *gin.Engine {
	service := gin.Default()
	service.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "medastra")
	})
	service.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Live")
	})
	service.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, "Not found")
	})

	router := service.Group("/api/v1/")
	userRouter := router.Group("/user/")
	doctorRouter := router.Group("/doctor/")

	doctorRouter.POST("/search", controller.GetDoctors)
	doctorRouter.GET("/bookings/:id", controller.DoctorBookings)
	doctorRouter.POST("/login", controller.Login)
	doctorRouter.POST("/create", controller.CreateAccount)

	userRouter.POST("/login", controller.Login)
	userRouter.POST("/create", controller.CreateAccount)

	userRouter.POST("/userprofile/", controller.CreateUserProfile)
	userRouter.PUT("/userprofile/:username", controller.UpdateUserProfile)
	userRouter.GET("/userprofile/:username", controller.FindUserByUsername)
	userRouter.GET("/slot/:slotId/:username", controller.BookSlot)
	userRouter.GET("/bookings/:username", controller.UserBookings)

	return service
}
