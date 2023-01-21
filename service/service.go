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
