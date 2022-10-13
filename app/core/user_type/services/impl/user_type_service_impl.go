package impl

import (
	"service-api/app/core/user_type/entities"
	"service-api/app/core/user_type/repositories"
)

type userTypeService struct {
	repositories repositories.UserTypeRepository
}

func NewUserTypeService(repo repositories.UserTypeRepository) *userTypeService {
	return &userTypeService{
		repositories: repo,
	}
}

func (service *userTypeService) GetAll() ([]entities.UserType, error) {
	userTypes, err := service.repositories.FindAll()
	if err != nil {
		return nil, err
	}

	return userTypes, nil
}
