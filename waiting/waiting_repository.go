package waiting

type Repository struct {
	waitingUsers []Waiting
}

func (r *Repository) Append(waitingUser *Waiting) []Waiting {
	r.waitingUsers = append(r.waitingUsers, *waitingUser)
	return r.waitingUsers
}

func (r *Repository) Remove(waitingUser *Waiting) []Waiting {

	for idx, v := range r.waitingUsers {
		if v.user.ID == waitingUser.user.ID {
			r.waitingUsers = append(r.waitingUsers[0:idx], r.waitingUsers[idx+1:]...)
			return r.waitingUsers
		}
	}
	return r.waitingUsers
}

func (r *Repository) GetMatches() [][]Waiting {
	return [][]Waiting{r.waitingUsers}
}

func NewRepository() *Repository {
	return &Repository{}
}

func InitRepository() *Repository {
	return NewRepository()
}
