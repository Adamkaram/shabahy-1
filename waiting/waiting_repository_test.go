package waiting

import (
	"github.com/ElegantSoft/shabahy/users"
	"testing"
)

func TestWaitingRepo(t *testing.T) {
	repo := InitRepository()
	t.Run("Test adding waiting", func(t *testing.T) {
		repo.Append(&Waiting{user: users.User{ID: 1, Name: "Ahmed"}})
		repo.Append(&Waiting{user: users.User{ID: 2, Name: "Bebo"}})
		if repo.waitingUsers[0].user.ID != 1 {
			t.Error("User not appended")
		}
		if repo.waitingUsers[1].user.ID != 2 {
			t.Error("User not appended")
		}
		if len(repo.waitingUsers) != 2 {
			t.Error("Issue in append user")
		}
	})

	t.Run("Test Removing user", func(t *testing.T) {
		repo.Remove(&Waiting{user: users.User{ID: 1, Name: "Ahmed"}})

		if repo.waitingUsers[0].user.ID != 2 {
			t.Error("User not removed")
		}

		if len(repo.waitingUsers) != 1 {
			t.Error("Issue in remove user")
		}
	})


}
