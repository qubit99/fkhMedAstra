package models

import (
	"time"
)

type UserProfile struct {
	Username   string    `gorm:"type:varchar(255);primary_key" json:"username"`
	Height     int       `gorm:"type:int" json:"height"`
	Weight     int       `gorm:"type:int" json:"weight"`
	Gender     string    `gorm:"type:varchar(255)" json:"gender"`
	DOB        time.Time `gorm:"type:date" json:"dob"`
	Name       string    `gorm:"type:varchar(255)" json:"name"`
	City       string    `gorm:"type:varchar(255)" json:"city"`
	BloodGroup string    `gorm:"type:varchar(4)" json:"blood_group"`
}
type User struct {
	Username string `gorm:"type:varchar(255);primary_key" json:"username"`
	Password string `gorm:"type:varchar(255)" json:"password"`
}
