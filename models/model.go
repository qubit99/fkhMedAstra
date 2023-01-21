package models

type Slot struct {
	SlotId   int    `gorm:"type:int;primary_key" json:"slot_id"`
	DoctorId int    `gorm:"type:int" json:"doctor_id"`
	Interval string `gorm:"type:string" json:"interval"`
	Date     string `gorm:"type:date" json:"date"`
	Capacity int    `gorm:"type:int" json:"capacity"`
}

type Booking struct {
	SlotId   int
	Username string
	DoctorId int
}

type UserBookingResponse struct {
	Slot   *Slot   `json:"slot"`
	Doctor *Doctor `json:"doctor"`
}
type DoctorBookingResponse struct {
	Slot *Slot        `json:"slot"`
	User *UserProfile `json:"user"`
}

/*
GET /getUserbyId/:id
POST /signup/ {username: , password: } --> userId
POST /createProfile/ {userprofile} --> userprofile
PUT /updateProfile/:username {userprofile} --> userprofile
GET /searchDoctor/ {doctorsearch request} -> []Doctors
GET /getSlots/ {slotsearch } --> SlotSearchResponse[]
PUT /bookSlot/:id?userId=userId -->
GET /getBookings/:userId -> []Slots








*/
