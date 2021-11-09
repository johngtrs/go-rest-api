package movie

type MovieService interface {
	ListService() ([]Movie, error)
	readByIdService(id string) (Movie, error)
	createService(movie Movie) (int64, error)
}

type Service struct {
	repository MovieRepository
}

func NewMovieService(repository MovieRepository) *Service {
	return &Service{repository: repository}
}

func (s *Service) ListService() ([]Movie, error) {
	return s.repository.FindAll()
}

func (s *Service) readByIdService(id string) (Movie, error) {
	return s.repository.FindFirst(id)
}

func (s *Service) createService(movie Movie) (int64, error) {
	return s.repository.AddMovie(movie)
}
