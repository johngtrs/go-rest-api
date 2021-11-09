package album

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AlbumRepository interface {
	FindAll() ([]Album, error)
	FindFirst(id string) (Album, error)
	FindByArtist(name string) ([]Album, error)
	AddAlbum(album Album) (int64, error)
}

type Repository struct {
	db *sqlx.DB
}

const table = "album"

func NewAlbumRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindAll() ([]Album, error) {
	albums := []Album{}

	err := r.db.Select(&albums, "SELECT * FROM "+table)
	if err != nil {
		return nil, fmt.Errorf("Album.FindAll: %v", err)
	}

	return albums, nil
}

func (r *Repository) FindFirst(id string) (Album, error) {
	var album Album

	err := r.db.Get(&album, "SELECT * FROM "+table+" WHERE id = ?", id)
	if err != nil {
		return album, fmt.Errorf("Album not found")
	}

	return album, nil
}

func (r *Repository) FindByArtist(name string) ([]Album, error) {
	albums := []Album{}

	err := r.db.Select(&albums, "SELECT * FROM "+table+" WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("FindByArtist %q: %v", name, err)
	}

	return albums, nil
}

func (r *Repository) AddAlbum(album Album) (int64, error) {
	result, err := r.db.Exec("INSERT INTO "+table+" (title, artist, price) VALUES (?, ?, ?)", album.Title, album.Artist, album.Price)
	if err != nil {
		return 0, fmt.Errorf("AddAlbum: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("AddAlbum: %v", err)
	}

	return id, nil
}
