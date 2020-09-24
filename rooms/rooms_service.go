package rooms

type Service struct {
	repo Repository
}


func (s *Service) Paginate() (error, interface{}) {
	return s.repo.Paginate()
}

func (s *Service) Find(id uint) (error, interface{}) {
	return s.repo.Find(id)
}

func (s *Service) Create(item *Room) (error, interface{}) {
	return s.repo.Create(item)
}

func (s *Service) Update(item *Room, id uint) error {
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