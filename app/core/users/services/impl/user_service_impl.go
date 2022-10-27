package impl

import (
	"errors"
	userTypeRepositories "service-api/app/core/user_type/repositories"
	"service-api/app/core/users/entities"
	"service-api/app/core/users/repositories"
	"service-api/app/util/users"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo     repositories.UserRepository
	userTypeRepo userTypeRepositories.UserTypeRepository
}

func NewUserService(userRepo repositories.UserRepository, userTypeRepo userTypeRepositories.UserTypeRepository) *userService {
	return &userService{
		userRepo:     userRepo,
		userTypeRepo: userTypeRepo,
	}
}

func (service *userService) FindAll() ([]entities.User, error) {
	users, err := service.userRepo.FindAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (service *userService) RegisterUser(user users.RegisterUserInput) (entities.User, error) {
	userData := entities.User{}
	userData.Name = user.Name
	userData.Username = user.Username
	userData.Email = user.Email
	userData.Password = user.Password

	// hashing
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.MinCost)
	if err != nil {
		return userData, err
	}
	userData.Password = string(passwordHash)

	// fetch user type
	userType, err := service.userTypeRepo.FindUserTypeById(user.UserTypeId)
	if err != nil {
		return userData, err
	}
	userData.UserType = userType

	newUser, err := service.userRepo.Save(userData)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (service *userService) LoginUser(userInput users.LoginUserInput) (entities.User, error) {
	email := userInput.Email
	password := userInput.Password

	user, err := service.userRepo.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("user not found")
	}

	// compare hasing password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}

func (handler *userService) GetUserByID(id uint) (entities.User, error) {
	user, err := handler.userRepo.FindByID(id)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("user not found")
	}

	return user, nil
}
