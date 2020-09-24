package rooms

import (
	"github.com/ElegantSoft/shabahy/users"
	"math/rand"
	"strconv"
)

type Service struct {
	repo Repository
	userService users.Service
}


func (s *Service) Paginate() (error, interface{}) {
	return s.repo.Paginate()
}

func (s *Service) Find(id uint) (error, interface{}) {
	return s.repo.Find(id)
}

func (s *Service) Create(usersFromRequest []uint) (error, interface{}) {
	hash := s.GenerateHash()
	usersToAppend := users.GetUsersFromIds(usersFromRequest)
	return s.repo.Create(hash, &usersToAppend)
}

func (s *Service) GenerateHash() string {
	randomNumber := rand.Int()
	return strconv.Itoa(randomNumber)
}

func (s *Service) Update(item *Room, id uint) error {
	return s.repo.Update(item, id)
}

func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}



func NewService(repository *Repository ) *Service  {
	return &Service{
		repo: *repository,
	}
}