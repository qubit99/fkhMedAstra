package main

import (
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
	/*
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
		/*
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
				err = repo.Db.Table(repository.DoctorTable).Create(&doctor).Omit("id").Error
				if err != nil {
					log.Println(err)
				}
			}

			for i := 100; i < 1000; i++ {
				username := usercreds[i].Username
				user := models.UserProfile{}
				err := faker.FakeData(&user)
				if err != nil {
					log.Println(err)
				}
				user.Username = username
				err = repo.SaveProfile(&user)
			}
	*/

	doctors := make([]*models.Doctor, 0)
	err := repo.Db.Table(repository.DoctorTable).Find(&doctors).Error
	if err != nil {
		log.Println(err)
		return
	}
	dates := []string{"2022-01-25", "2022-01-26", "2022-01-27", "2022-01-28"}
	intervals := []string{"9AM-12PM", "2PM-5PM", "6PM-9PM"}
	defaultCapacity := 4
	for _, doctor := range doctors {
		//create slots
		for _, date := range dates {
			for _, interval := range intervals {
				slot := models.Slot{
					DoctorId: doctor.Id,
					Capacity: defaultCapacity,
					Interval: interval,
					Date:     date,
				}
				err := repo.Db.Table(repository.SlotTable).Create(&slot).Error
				if err != nil {
					log.Println(err)
					continue
				}
			}
		}
	}

}
