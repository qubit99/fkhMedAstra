package models

type Doctor struct {
	Id         int    `gorm:"type:int;primary_key" json:"id"`
	Username   string `gorm:"type:varchar(255)" json:"username"`
	Name       string `gorm:"type:varchar(255)" json:"name" faker:"name"`
	Speciality string `gorm:"type:varchar(255)" json:"speciality" faker:"oneof: ENT, Orthopaedic, Gynaecology, Neurosurgery, Oncology, Urology, Dermatology, Psychiatry, Pediatrics"`
	Rating     int    `gorm:"type:int" json:"rating" faker:"oneof: 5, 4, 3, 2, 1"`
	Fee        int    `gorm:"type:int" json:"fee" faker:"oneof: 500, 1000, 1500, 2000, 2500"`
	City       string `gorm:"type:varchar(255)" json:"city" faker:"oneof: Bengaluru, Mumbai, Delhi, Gurgaon, Hyderabad"`
}

type DoctorResponse struct {
	/*
		Id         int     `json:"id"`
		Username   string  `json:"username"`
		Name       string  `json:"name"`
		Rating     int     `json:"rating"`
		Speciality string  `json:"speciality"`
		Fee        int     `json:"fee"`
		City       string  `json:"city"`

	*/
	*Doctor
	Slots []*Slot `json:"slots"`
}

type DoctorSearchRequest struct {
	Ids          []int    `json:"ids"`
	Names        []string `json:"names"`
	Specialities []string `json:"specialities"`
	Ratings      []int    `json:"ratings"`
	Cities       []string `json:"cities"`
	SortOrder    string   `json:"sort_order"`
	SortBy       string   `json:"sort_by"`
	Limit        int      `json:"limit"`
	Offset       int      `json:"offset"`
}
