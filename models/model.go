package models

import "time"

type Slots struct {
	SlotId   int
	DoctorId int
	Interval string
	Date     time.Time
	Capacity int
}

type DiagnosticsTest struct {
	//report
}

type BookDiagnosticTest struct {
	Type   string
	Fee    int
	Date   time.Time
	UserId int
}

type Booking struct {
	SlotId   int
	Username int
}

type SlotSearchRequest struct {
	Specialities []string
	City         string
	orderby      string
}
type SlotSearchResponse struct {
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
