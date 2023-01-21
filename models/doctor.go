package models

type Doctor struct {
	Id         int    `gorm:"type:int"`
	Name       string `gorm:"type:varchar(255)"`
	Speciality string `gorm:"type:varchar(255)"`
	Rating     int    `gorm:"type:int"`
	Fee        int    `gorm:"type:int"`
	City       string `gorm:"type:varchar(255)"`
}

type DoctorResponse struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Speciality string `json:"speciality"`
	Rating     int    `json:"rating"`
	Fee        int    `json:"fee"`
	City       string `json:"city"`
}

func (d *Doctor) GetResponseFromEntity() *DoctorResponse {
	return &DoctorResponse{
		Id:         d.Id,
		Fee:        d.Fee,
		Name:       d.Name,
		Rating:     d.Rating,
		City:       d.City,
		Speciality: d.Speciality,
	}
}

/*
func (d *DoctorRequest) GetRequestFromEntity() *Doctor {
	return &Doctor{
		Id:         d.Id,
		Fee:        d.Fee,
		Name:       d.Name,
		Rating:     d.Rating,
		City:       d.City,
		Speciality: d.Speciality,
	}
}
*/
