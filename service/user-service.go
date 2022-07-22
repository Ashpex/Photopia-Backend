package service

import (
	"github.com/mashingan/smapping"
	"gitlab.zalopay.vn/top/intern/vybnt/gallery-backend/gallery/dto"
	"gitlab.zalopay.vn/top/intern/vybnt/gallery-backend/gallery/entity"
	"gitlab.zalopay.vn/top/intern/vybnt/gallery-backend/gallery/repository"
	"log"
)

// UserService is a contract about something that this service can do
type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *userService) Profile(userID string) entity.User {
	return service.userRepository.ProfileUser(userID)
}
