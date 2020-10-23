package waiting

type Repository interface {
	append(waitingUser *Waiting) []Waiting
}

type RepositoryImpl struct {
	waitingUsers []Waiting
}

func (r *RepositoryImpl) append(waitingUser *Waiting) []Waiting {
	return append(r.waitingUsers, *waitingUser)
}

func NewRepository() Repository {
	return &RepositoryImpl{}
}

func InitRepository() Repository {
	return NewRepository()
}