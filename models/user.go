package models

type UserProfile struct {
	Username   string `gorm:"type:varchar(255);primary_key" json:"username"`
	Height     int    `gorm:"type:int" json:"height" faker:"oneof: 180, 190, 200, 140, 130"`
	Weight     int    `gorm:"type:int" json:"weight" faker:"oneof: 60, 70, 80, 90, 100"`
	Gender     string `gorm:"type:varchar(255)" json:"gender" faker:"oneof: Male, Female, Other"`
	DOB        string `gorm:"type:varchar(255)" json:"dob" faker:"date"`
	Name       string `gorm:"type:varchar(255)" json:"name" faker:"name"`
	City       string `gorm:"type:varchar(255)" json:"city" faker:"oneof: Bengaluru, Mumbai, Delhi, Gurgaon, Hyderabad"`
	BloodGroup string `gorm:"type:varchar(4)" json:"blood_group" faker:"oneof: A+, B+, O+, O-, B-, A-"`
}
type User struct {
	Username string `gorm:"type:varchar(255);primary_key" json:"username" faker:"username"`
	Password string `gorm:"type:varchar(255)" json:"password" faker:"oneof: password, 123, 321, 546"`
}
