package movie

import (
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type MovieRepository interface {
	FindAll() ([]Movie, error)
	FindFirst(id string) (Movie, error)
	MostRentedList(year string, limit int) ([]Movie, error)
	MostRented(year string) (Movie, error)
	FindBestAuthor() (map[string]string, error)
	FindByTitle(title string) ([]Movie, error)
	AddMovie(movie Movie) (int64, error)
	IncrementRentedNumber(title string, year string) error
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

func (r *Repository) MostRentedList(year string, limit int) ([]Movie, error) {
	movies := []Movie{}
	var err error

	if year != "" {
		sql := "SELECT * FROM " + table + " WHERE year = ? ORDER BY rent_number DESC LIMIT " + strconv.Itoa(limit)
		err = r.db.Select(&movies, sql, year)
	} else {
		sql := "SELECT * FROM " + table + " ORDER BY rent_number DESC LIMIT " + strconv.Itoa(limit)
		err = r.db.Select(&movies, sql)
	}

	if err != nil {
		return nil, fmt.Errorf("Movie.MostRentedList: %v", err)
	}

	return movies, nil
}

func (r *Repository) MostRented(year string) (Movie, error) {
	var movie Movie
	var err error

	if year != "" {
		sql := "SELECT * FROM " + table + " WHERE year = ? ORDER BY rent_number DESC"
		err = r.db.Get(&movie, sql, year)
	} else {
		sql := "SELECT * FROM " + table + " ORDER BY rent_number DESC"
		err = r.db.Get(&movie, sql)
	}

	if err != nil {
		return movie, fmt.Errorf("Movie.MostRented: %v", err)
	}

	return movie, nil
}

func (r *Repository) FindBestAuthor() (map[string]string, error) {
	var movie Movie
	var err error

	sql := "SELECT author FROM " + table + " ORDER BY rent_number DESC"
	err = r.db.Get(&movie, sql)

	if err != nil {
		return nil, fmt.Errorf("Movie.FindBestAuthor: %v", err)
	}

	data := make(map[string]string)
	data["author"] = movie.Author

	return data, nil
}

func (r *Repository) FindByTitle(title string) ([]Movie, error) {
	movies := []Movie{}

	err := r.db.Select(&movies, "SELECT * FROM "+table+" WHERE title LIKE ?", "%"+title+"%")
	if err != nil {
		return nil, fmt.Errorf("Movie.FindByTitle %q: %v", title, err)
	}

	return movies, nil
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

func (r *Repository) IncrementRentedNumber(title string, year string) error {
	var movie Movie

	errSelect := r.db.Get(&movie, "SELECT * FROM movie WHERE title=? AND year=?", title, year)
	if errSelect != nil {
		return fmt.Errorf("Movie not found")
	}

	sql := "UPDATE " + table + " SET rent_number = rent_number + 1 WHERE title=? AND year=?"

	_, errUpdate := r.db.Exec(sql, title, year)
	if errUpdate != nil {
		return fmt.Errorf("IncrementRentedNumber: %v", errUpdate)
	}

	return nil
}
