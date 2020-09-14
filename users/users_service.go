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
	if err, createdUser := s.repo.Create(user); err != nil {
		return err, nil
	} else {
		return nil, createdUser
	}
}

func (s Service) Login(data *LoginUserDTO) (error, string) {
	if err, user := s.repo.FindUserByIdAndPassword(data); err == nil {
		atClaims := jwt.MapClaims{}
		atClaims["user_id"] = user.ID
		at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
		token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
		if err != nil {
			return err, ""
		}
		return err, token
	} else {
		return err, ""
	}
}

func NewService(repository *Repository) *Service {
	return &Service{
		repo: *repository,
	}
}
