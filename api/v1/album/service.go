package album

import "github.com/johngtrs/go-rest-api/model"

type AlbumService interface {
	ListService() ([]model.Album, error)
	readByIdService(id string) (model.Album, error)
	listByArtistService(name string) ([]model.Album, error)
	createService(album model.Album) (int64, error)
}

type Service struct {
	repository AlbumRepository
}

func NewAlbumService(repository AlbumRepository) *Service {
	return &Service{repository: repository}
}

func (s *Service) ListService() ([]model.Album, error) {
	return s.repository.FindAll()
}

func (s *Service) readByIdService(id string) (model.Album, error) {
	return s.repository.FindFirst(id)
}

func (s *Service) listByArtistService(name string) ([]model.Album, error) {
	return s.repository.FindByArtist(name)
}

func (s *Service) createService(album model.Album) (int64, error) {
	return s.repository.AddAlbum(album)
}
