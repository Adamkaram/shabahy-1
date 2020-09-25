package rooms

import (
	"fmt"
	"github.com/ElegantSoft/shabahy/common"
	"log"
	"math/rand"
	"strconv"
)

type Service struct {
	repo Repository
}

func (s *Service) Paginate() (error, interface{}) {
	return s.repo.Paginate()
}

func (s *Service) Find(id uint) (error, interface{}) {
	return s.repo.Find(id)
}

func (s *Service) GetUsersFromIds(ids []uint) []User {
	var usersToAppend = make([]User, 0)
	for i := 0; i < len(ids); i++ {
		usersToAppend = append(usersToAppend, User{ID: ids[i]})
	}
	return usersToAppend
}

func (s *Service) Create(usersFromRequest []uint) (error, *Room) {
	hash := s.GenerateHash()
	usersToAppend := s.GetUsersFromIds(usersFromRequest)
	uniqueIds := common.Unique(usersFromRequest)
	if len(uniqueIds) < 2 {
		return fmt.Errorf("room should have more than user"), nil
	}
	err, exists := s.repo.FindRoomWithUsersIds(uniqueIds)
	if err != nil {
		log.Println("err1", err)
		return err, nil
	}
	if exists {
		return fmt.Errorf("room already exists"), nil
	}
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

func NewService(repository *Repository) *Service {
	return &Service{
		repo: *repository,
	}
}
