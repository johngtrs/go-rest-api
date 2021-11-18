package album

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/johngtrs/go-rest-api/database"
	"github.com/johngtrs/go-rest-api/glogger"
	"github.com/johngtrs/go-rest-api/httperror"
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
	var albums []Album
	builder := database.NewQueryBuilder(r.db, &albums)

	err := builder.Select("*").From(table, "").Exec()
	if err != nil {
		glogger.Log("Album.FindAll", err.Error())
		return nil, httperror.ErrInternalServerError
	}

	return albums, nil
}

func (r *Repository) FindFirst(id string) (Album, error) {
	var album Album
	builder := database.NewQueryBuilder(r.db, &album)

	err := builder.Select("*").From(table, "").Where("id = ?", id).ExecOne()
	if err != nil {
		if err == sql.ErrNoRows {
			return album, httperror.ErrNotFound
		}
		glogger.Log("Album.FindFirst", err.Error())
		return album, httperror.ErrInternalServerError
	}

	return album, nil
}

func (r *Repository) FindByArtist(name string) ([]Album, error) {
	albums := []Album{}
	builder := database.NewQueryBuilder(r.db, &albums)

	err := builder.Select("*").From(table, "").Where("artist = ?", name).Exec()
	if err != nil {
		glogger.Log("Album.FindByArtist", err.Error())
		return nil, httperror.ErrInternalServerError
	}

	return albums, nil
}

func (r *Repository) AddAlbum(album Album) (int64, error) {
	result, err := r.db.Exec("INSERT INTO "+table+" (title, artist, price) VALUES (?, ?, ?)", album.Title, album.Artist, album.Price)
	if err != nil {
		glogger.Log("Album.AddAlbum", err.Error())
		return 0, httperror.ErrInternalServerError
	}

	id, err := result.LastInsertId()
	if err != nil {
		glogger.Log("Album.AddAlbum", err.Error())
		return 0, httperror.ErrInternalServerError
	}

	return id, nil
}
