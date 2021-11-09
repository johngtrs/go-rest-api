package movie

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type MovieRepository interface {
	FindAll() ([]Movie, error)
	FindFirst(id string) (Movie, error)
	AddMovie(movie Movie) (int64, error)
}

type Repository struct {
	db *sqlx.DB
}

const table = "movie"

func NewMovieRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindAll() ([]Movie, error) {
	movies := []Movie{}

	err := r.db.Select(&movies, "SELECT * FROM "+table)
	if err != nil {
		return nil, fmt.Errorf("Movie.FindAll: %v", err)
	}

	return movies, nil
}

func (r *Repository) FindFirst(id string) (Movie, error) {
	var movie Movie

	err := r.db.Get(&movie, "SELECT * FROM "+table+" WHERE id = ?", id)
	if err != nil {
		return movie, fmt.Errorf("Movie not found")
	}

	return movie, nil
}

func (r *Repository) AddMovie(movie Movie) (int64, error) {
	sql := "INSERT INTO " + table +
		" (year, rent_number, title, author, editor, `index`, bib, ref, cat_1, cat_2) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := r.db.Exec(sql, movie.Year, movie.RentNumber, movie.Title, movie.Author,
		movie.Editor, movie.Index, movie.Bib, movie.Ref, movie.Cat1, movie.Cat2)
	if err != nil {
		return 0, fmt.Errorf("AddMovie: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("AddMovie: %v", err)
	}

	return id, nil
}
