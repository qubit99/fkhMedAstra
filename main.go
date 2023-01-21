package main

import (
	"log"
	"medastra/config"
	"medastra/controller"
	"medastra/models"
	"medastra/repository"
	"medastra/routes"
	"medastra/service"
	"net/http"
	"os"
)

func main() {
	db := config.DatabaseConnection()
	db.Table(repository.UserTable).AutoMigrate(&models.UserProfile{})
	db.Table(repository.LoginTable).AutoMigrate(&models.User{})
	repo := repository.NewUserRepository(db)
	svc := service.NewServiceImpl(repo)
	ctrl := controller.NewController(svc)
	router := routes.NewRouter(ctrl)

	server := http.Server{
		Addr:    ":8888",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
