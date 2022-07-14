package service

import (
	"example.com/gallery/dto"
	"example.com/gallery/entity"
	"example.com/gallery/repository"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// AuthService ias a contract about something that this service can do
type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.RegisterDTO) entity.User
	FindUserByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

// NewAuthService function creates a new instance of AuthService
func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRepo,
	}
}

func (service *authService) VerifyCredential(email string, password string) interface{} {
	res := service.userRepository.VerifyCredential(email, password)
	if v, ok := res.(entity.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
			log.Println("User is valid")
		} else {
			log.Println("User is invalid")
			return false
		}
		return false

	}
	return false
}

func (service *authService) CreateUser(user dto.RegisterDTO) entity.User {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Println(err)
		log.Fatalf("Failed to create user")
	}
	return service.userRepository.InsertUser(userToCreate)
}

func (service *authService) FindUserByEmail(email string) entity.User {
	return service.userRepository.FindUserByEmail(email)
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparePassword(hashedPassword string, plainPassword []byte) bool {
	byteHash := []byte(hashedPassword)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
