package movie

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/johngtrs/go-rest-api/utils"
)

type MovieRepository interface {
	FindAll() ([]Movie, error)
	FindFirst(id string) (Movie, error)
	MostRentedList(year string, limit int) ([]Movie, error)
	MostRented(year string) (Movie, error)
	FindBestAuthor() (string, error)
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
		utils.Glogger("Movie.FindAll", err.Error())
		return nil, utils.ErrInternalServerError
	}

	return movies, nil
}

func (r *Repository) FindFirst(id string) (Movie, error) {
	var movie Movie

	err := r.db.Get(&movie, "SELECT * FROM "+table+" WHERE id = ?", id)
	if err != nil {
		utils.Glogger("Movie.FindFirst", err.Error())
		return movie, utils.ErrNotFound
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
		utils.Glogger("Movie.MostRentedList", err.Error())
		return nil, utils.ErrInternalServerError
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
		utils.Glogger("Movie.MostRented", err.Error())
		return movie, utils.ErrNotFound
	}

	return movie, nil
}

func (r *Repository) FindBestAuthor() (string, error) {
	var movie Movie
	var err error

	sql := "SELECT author FROM " + table + " ORDER BY rent_number DESC"
	err = r.db.Get(&movie, sql)

	if err != nil {
		utils.Glogger("Movie.FindBestAuthor", err.Error())
		return "", utils.ErrNotFound
	}

	return movie.Author, nil
}

func (r *Repository) FindByTitle(title string) ([]Movie, error) {
	movies := []Movie{}

	err := r.db.Select(&movies, "SELECT * FROM "+table+" WHERE title LIKE ?", "%"+title+"%")
	if err != nil {
		utils.Glogger("Movie.FindByTitle", err.Error())
		return nil, utils.ErrInternalServerError
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
		utils.Glogger("Movie.AddMovie", err.Error())
		return 0, utils.ErrInternalServerError
	}

	id, err := result.LastInsertId()
	if err != nil {
		utils.Glogger("Movie.AddMovie", err.Error())
		return 0, utils.ErrInternalServerError
	}

	return id, nil
}

func (r *Repository) IncrementRentedNumber(title string, year string) error {
	var movie Movie

	err := r.db.Get(&movie, "SELECT * FROM movie WHERE title=? AND year=?", title, year)
	if err != nil {
		utils.Glogger("Movie.IncrementRentedNumber", err.Error())
		return utils.ErrNotFound
	}

	sql := "UPDATE " + table + " SET rent_number = rent_number + 1 WHERE title=? AND year=?"

	_, err = r.db.Exec(sql, title, year)
	if err != nil {
		utils.Glogger("Movie.IncrementRentedNumber", err.Error())
		return utils.ErrInternalServerError
	}

	return nil
}
