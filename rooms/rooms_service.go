package rooms

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ElegantSoft/shabahy/common"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type Service struct {
	repo Repository
}

func (s *Service) paginate() (error, interface{}) {
	return s.repo.paginate()
}

func (s *Service) find(id uint, userId uint) (error, *Room) {
	roomUsers := s.repo.getUsers(id)
	if !common.Contains(userId, roomUsers) {
		return errors.New("you can't get messages from this room"), nil
	}
	return s.repo.find(id)
}

func (s *Service) getUsersFromIds(ids []uint) []User {
	var usersToAppend = make([]User, 0)
	for i := 0; i < len(ids); i++ {
		usersToAppend = append(usersToAppend, User{ID: ids[i]})
	}
	return usersToAppend
}

func (s *Service) create(usersFromRequest []uint) (error, *Room) {
	hash := s.generateHash()
	usersToAppend := s.getUsersFromIds(usersFromRequest)
	uniqueIds := common.Unique(usersFromRequest)
	if len(uniqueIds) < 2 {
		return fmt.Errorf("room should have more than user"), nil
	}
	err, exists := s.repo.findRoomWithUsersIds(uniqueIds)
	if err != nil {
		log.Println("err1", err)
		return err, nil
	}
	if exists {
		return fmt.Errorf("room already exists"), nil
	}
	return s.repo.create(hash, &usersToAppend)
}

func (s *Service) AppendMessage(roomId uint, message *Message, userId uint) error {
	room := &Room{
		ID: roomId,
	}
	roomUsers := s.repo.getUsers(roomId)
	log.Println("userid", userId)
	if !common.Contains(userId, roomUsers) {
		return errors.New("you can't send message to this room")
	}
	message.UserID = userId
	return s.repo.appendMessage(room, message)
}

func (s *Service) generateHash() string {
	randomNumber := strconv.Itoa(rand.Int()) + time.Now().String()
	sha := sha256.New()
	sha.Write([]byte(randomNumber))
	hash := hex.EncodeToString(sha.Sum(nil))
	return hash
}

func (s *Service) update(item *Room, id uint) error {
	return s.repo.update(item, id)
}

func (s *Service) delete(id uint) error {
	return s.repo.delete(id)
}

func NewService(repository *Repository) *Service {
	return &Service{
		repo: *repository,
	}
}
