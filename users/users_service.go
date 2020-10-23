package users

import (
	"github.com/ElegantSoft/shabahy/services"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Create(user *User) (string, error)
	Login(data *LoginUserDTO) (string, *User, error)
}

type serviceImpl struct {
	repo Repository
	jwtService services.JWTService
}


func (s serviceImpl) Create(user *User) (string, error) {
	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	user.Password = string(password)
	err, _ := s.repo.Create(user)
	if err != nil {
		return "", err
	}
	token, err := s.jwtService.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s serviceImpl) Login(data *LoginUserDTO) (string, *User, error) {
	err, user := s.repo.FindUserByIdAndPassword(data)
	if err != nil {
		return "", nil, err
	}
	token, err := s.jwtService.GenerateToken(user.ID)
	if err != nil {
		return "", nil, err
	}
	return token, user, nil
}



func NewService(repository *Repository, jwtService *services.JWTService) Service {
	return &serviceImpl{
		repo: *repository,
		jwtService: *jwtService,
	}
}

func InitUserService() Service {
	jwtService := services.NewJWTService()
	repository := *NewRepository()
	return NewService(&repository, &jwtService)
}