package models

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AlbumRepository struct {
	*sqlx.DB
}

type Album struct {
	ID     uint32  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

const modelName = "album"

func (r *AlbumRepository) FindAll() ([]Album, error) {
	albums := []Album{}

	err := r.DB.Select(&albums, "SELECT * FROM "+modelName)
	if err != nil {
		return nil, fmt.Errorf("Album.FindAll: %v", err)
	}

	return albums, nil
}

// albumByID queries for the album with the specified ID.
func (r *AlbumRepository) FindFirst(id uint64) (Album, error) {
	var alb Album

	err := r.DB.Get(&alb, "SELECT * FROM album WHERE id = ?", id)
	if err != nil {
		return alb, fmt.Errorf("Album not found")
	}
	return alb, nil
}

// albumsByArtist queries for albums that have the specified artist name.
func (r *AlbumRepository) FindByArtist(name string) ([]Album, error) {
	// An albums slice to hold data from returned rows.
	albums := []Album{}

	err := r.DB.Select(&albums, "SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("GetAlbumsByArtist %q: %v", name, err)
	}

	return albums, nil
}

// AddAlbum adds the specified album to the database,
// returning the album ID of the new entry
func (r *AlbumRepository) AddAlbum(alb Album) (int64, error) {
	result, err := r.DB.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("AddAlbum: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("AddAlbum: %v", err)
	}

	return id, nil
}
