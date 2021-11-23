package movie

import "github.com/johngtrs/go-rest-api/model"

type MovieService interface {
	ListService() ([]model.Movie, error)
	readByIdService(id string) (model.Movie, error)
	Top100Service() ([]model.Movie, error)
	Top100YearService(year string) ([]model.Movie, error)
	MostRentedService() (model.Movie, error)
	MostRentedYearService(year string) (model.Movie, error)
	BestAuthorService() (map[string]string, error)
	SearchByTitleService(title string) ([]model.Movie, error)
	createService(movie model.Movie) (int64, error)
	IncrementRentedNumberService(title string, year string) error
}

type Service struct {
	repository MovieRepository
}

func NewMovieService(repository MovieRepository) *Service {
	return &Service{repository: repository}
}

func (s *Service) ListService() ([]model.Movie, error) {
	return s.repository.FindAll()
}

func (s *Service) readByIdService(id string) (model.Movie, error) {
	return s.repository.FindFirst(id)
}

func (s *Service) Top100Service() ([]model.Movie, error) {
	return s.repository.MostRentedList("", 100)
}

func (s *Service) Top100YearService(year string) ([]model.Movie, error) {
	return s.repository.MostRentedList(year, 100)
}

func (s *Service) MostRentedService() (model.Movie, error) {
	return s.repository.MostRented("")
}

func (s *Service) MostRentedYearService(year string) (model.Movie, error) {
	return s.repository.MostRented(year)
}

func (s *Service) BestAuthorService() (map[string]string, error) {
	author, err := s.repository.FindBestAuthor()
	if err != nil {
		return nil, err
	}

	data := make(map[string]string)
	data["author"] = author

	return data, nil
}

func (s *Service) SearchByTitleService(title string) ([]model.Movie, error) {
	return s.repository.FindByTitle(title)
}

func (s *Service) createService(movie model.Movie) (int64, error) {
	return s.repository.AddMovie(movie)
}

func (s *Service) IncrementRentedNumberService(title string, year string) error {
	return s.repository.IncrementRentedNumber(title, year)
}
