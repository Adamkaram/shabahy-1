package users

import "golang.org/x/crypto/bcrypt"

type Service struct {
	repo Repository
}

func (s Service) Create(user *User) (error, *User) {
	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 0)
	user.Password = string(password)
	if err, createdUser := s.repo.Create(user); err != nil {
		return err, nil
	} else {
		return nil, createdUser
	}
}

func (s Service) Login(id uint) (error, string) {
	if err, user := s.repo.FindUserById(id); err == nil {

	} else {
		return err, ""
	}
}

func NewService(repository *Repository) *Service {
	return &Service{
		repo: *repository,
	}
}
