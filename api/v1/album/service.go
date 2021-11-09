package album

type AlbumService interface {
	ListService() ([]Album, error)
	readByIdService(id string) (Album, error)
	listByArtistService(name string) ([]Album, error)
	createService(alb Album) (int64, error)
}

type Service struct {
	repository AlbumRepository
}

func NewAlbumService(repository AlbumRepository) *Service {
	return &Service{repository: repository}
}

func (s *Service) ListService() ([]Album, error) {
	return s.repository.FindAll()
}

func (s *Service) readByIdService(id string) (Album, error) {
	return s.repository.FindFirst(id)
}

func (s *Service) listByArtistService(name string) ([]Album, error) {
	return s.repository.FindByArtist(name)
}

func (s *Service) createService(album Album) (int64, error) {
	return s.repository.AddAlbum(album)
}
