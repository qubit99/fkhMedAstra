package main

import (
	"fmt"
	faker "github.com/bxcodec/faker/v4"
	"log"
	"medastra/config"
	"medastra/models"
	"medastra/repository"
)

func main() {
	db := config.DatabaseConnection()
	db.Table(repository.UserTable).AutoMigrate(&models.UserProfile{})
	db.Table(repository.LoginTable).AutoMigrate(&models.User{})
	db.Table(repository.BookingTable).AutoMigrate(&models.Booking{})
	db.Table(repository.SlotTable).AutoMigrate(&models.Slot{})
	db.Table(repository.DoctorTable).AutoMigrate(&models.Doctor{})
	repo := repository.NewUserRepository(db)
	//generate fake user creds, 100 doctors, 1000 patients

	for i := 0; i < 1000; i++ {
		user := models.User{}
		err := faker.FakeData(&user)
		if err != nil {
			log.Println(err)
		}
		user.Username = fmt.Sprintf("user%d", i+1)
		log.Println(user)
		err = repo.CreateLogin(&user)
		if err != nil {
			log.Println(err)
			continue
		}
	}
	usercreds := make([]*models.User, 0)
	err := repo.Db.Table(repository.LoginTable).Find(&usercreds).Error
	if err != nil {
		log.Println(err)
		return
	}
	//100 doctors, 1000 users
	//create 100 doctors
	for i := 0; i < 100; i++ {
		username := usercreds[i].Username
		doctor := models.Doctor{}
		err := faker.FakeData(&doctor)
		if err != nil {
			log.Println(err)
			continue
		}
		doctor.Username = username
		err = repo.Db.Table(repository.DoctorTable).Create(&doctor).Error
		if err != nil {
			log.Println(err)
		}
	}

	for i := 100; i < 1000; i++ {

	}

}
