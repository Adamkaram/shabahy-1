package messages

type Service struct {
	repo Repository
}


func (s *Service) Paginate() (error, interface{}) {
	return s.repo.Paginate()
}

func (s *Service) Find(id uint) (error, interface{}) {
	return s.repo.Find(id)
}

func (s *Service) Create(item *Message) (error, interface{}) {
	return s.repo.Create(item)
}

func (s *Service) Update(item *Message, id uint) error {
	return s.repo.Update(item, id)
}

func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}



func NewService(repository *Repository) *Service  {
	return &Service{
		repo: *repository,
	}
}