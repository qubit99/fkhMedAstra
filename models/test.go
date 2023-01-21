package models

import "time"

type Testbookings struct {
	Id       int       `gorm:"type:int;primary_key" json:"booking_id"`
	Username string    `gorm:"type:varchar(255)" json:"username"`
	Date     time.Time `gorm:"type:date" json:"date"`
	TestName string    `gorm:"type:varchar(255)" json:"testName"`
}

type TestResults struct {
	BookingId        int    `gorm:"type:int;primary_key" json:"booking_id"`
	TotalCholesterol string `gorm:"type:varchar(20)" json:"total_cholesterol"`
	Ldl              string `gorm:"type:varchar(20)" json:"ldl"`
	Hdl              string `gorm:"type:varchar(20)" json:"hdl"`
	Triglycerides    string `gorm:"type:varchar(20)" json:"triglycerides"`
	Glucose          string `gorm:"type:varchar(20)" json:"glucose"`
}

type Patients struct {
	Id       int    `gorm:"type:int;primary_key" json:"id"`
	Username string `gorm:"type:varchar(255)" json:"username"`
	Disease  string `gorm:"type:varchar(255)" json:"disease"`
}
