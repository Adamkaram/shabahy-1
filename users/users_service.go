package users

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"os"
)

type Service struct {
	repo Repository
}

func (s Service) Create(user *User) (string, error) {
	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	user.Password = string(password)
	err, _ := s.repo.Create(user)
	if err != nil {
		return "", err
	}
	token, err := s.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s Service) Login(data *LoginUserDTO) (string, *User, error) {
	err, user := s.repo.FindUserByIdAndPassword(data)
	if err != nil {
		return "", nil, err
	}
	token, err := s.GenerateToken(user.ID)
	if err != nil {
		return "", nil, err
	}
	return token, user, nil
}

func (s *Service) GenerateToken(userId uint) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userId
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	return at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

}

func NewService(repository *Repository) *Service {
	return &Service{
		repo: *repository,
	}
}
