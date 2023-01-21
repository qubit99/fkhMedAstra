package service

import (
	"log"
	"medastra/models"
	"medastra/repository"
)

type Service struct {
	repository *repository.RepositoryImpl
}

func NewServiceImpl(impl *repository.RepositoryImpl) Service {
	return Service{
		repository: impl,
	}
}
func (s Service) CreateUserProfile(request *models.UserProfile) error {
	log.Printf("%v", request)
	return s.repository.SaveProfile(request)
}
func (s Service) UpdateUserProfile(request *models.UserProfile) error {
	return s.repository.UpdateProfile(request)
}

func (s Service) GetUserProfile(userName string) (*models.UserProfile, error) {
	userProfile, err := s.repository.FindProfileByUserName(userName)
	if err != nil {
		return nil, err
	}
	return userProfile, nil
}

func (s Service) GetDoctors(req *models.DoctorSearchRequest) ([]*models.DoctorResponse, error) {
	doctors, err := s.repository.SearchDoctors(req)
	if err != nil {
		return nil, err
	}
	doctorsResponse := make([]*models.DoctorResponse, 0)
	for _, doctor := range doctors {
		slots, err := s.repository.GetSlotsByDoctor(doctor.Id)
		if err != nil {
			log.Printf("Error getting slots for doctor %d : %s\n", doctor.Id, err)
			continue
		}
		doctorResp := &models.DoctorResponse{
			doctor,
			slots,
		}

		doctorsResponse = append(doctorsResponse, doctorResp)
	}
	return doctorsResponse, nil
}

func (s Service) BookSlot(username string, slotId int) error {
	return s.repository.BookSlot(username, slotId)
}

func (s Service) UserBookings(username string) ([]*models.UserBookingResponse, error) {
	return s.repository.GetUserBookings(username)
}
func (s Service) DoctorBookings(doctorId int) ([]*models.DoctorBookingResponse, error) {
	return s.repository.GetDoctorBookings(doctorId)
}
func (s Service) CreateAccount(user *models.User) error {
	return s.repository.CreateLogin(user)
}
func (s Service) Login(username string, password string) error {
	return s.repository.CheckLogin(username, password)
}

func (s Service) CreateBooking(booking *models.Booking) error {
	return s.repository.CreateBooking(booking)
}

func (s Service) GetBookings(username string) ([]models.Booking, error) {
	bookings, err := s.repository.GetBookings(username)
	if err != nil {
		return nil, err
	}
	return bookings, nil
}
