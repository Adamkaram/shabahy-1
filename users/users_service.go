package users

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"os"
)

type Service struct {
	repo Repository
}

func (s Service) Create(user *User) (error, *User) {
	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	user.Password = string(password)
	return s.repo.Create(user)
}

func (s Service) Login(data *LoginUserDTO) (string, error) {
	err, user := s.repo.FindUserByIdAndPassword(data)
	if err != nil {
		return "", err
	}
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = user.ID
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	return at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
}

func NewService(repository *Repository) *Service {
	return &Service{
		repo: *repository,
	}
}
