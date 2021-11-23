package movie

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/johngtrs/go-rest-api/database"
	"github.com/johngtrs/go-rest-api/glogger"
	"github.com/johngtrs/go-rest-api/httperror"
	"github.com/johngtrs/go-rest-api/model"
)

type MovieRepository interface {
	FindAll() ([]model.Movie, error)
	FindFirst(id string) (model.Movie, error)
	MostRentedList(year string, limit int) ([]model.Movie, error)
	MostRented(year string) (model.Movie, error)
	FindBestAuthor() (string, error)
	FindByTitle(title string) ([]model.Movie, error)
	AddMovie(movie model.Movie) (int64, error)
	IncrementRentedNumber(title string, year string) error
}

type Repository struct {
	db *sqlx.DB
}

const table = "movie"

func NewMovieRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

// Get all movies.
func (r *Repository) FindAll() ([]model.Movie, error) {
	movies := []model.Movie{}
	builder := database.NewQueryBuilder(r.db, &movies)

	err := builder.Select("*").From(table, "").Exec()
	if err != nil {
		glogger.Log("Movie.FindAll", err.Error())
		return nil, httperror.ErrInternalServerError
	}

	return movies, nil
}

// Get the first movie with the requested id.
func (r *Repository) FindFirst(id string) (model.Movie, error) {
	var movie model.Movie
	builder := database.NewQueryBuilder(r.db, &movie)

	err := builder.Select("*").From(table, "").Where("id = ?", id).ExecOne()
	if err != nil {
		if err == sql.ErrNoRows {
			return movie, httperror.ErrNotFound
		}
		glogger.Log("Movie.FindFirst", err.Error())
		return movie, httperror.ErrInternalServerError
	}

	return movie, nil
}

// Get the most rented movies with the requested year.
// year parameter can be empty, in this case it will returns
// the most rented movies all of the times.
func (r *Repository) MostRentedList(year string, limit int) ([]model.Movie, error) {
	movies := []model.Movie{}
	var err error
	builder := database.NewQueryBuilder(r.db, &movies)

	if year != "" {
		err = builder.Select("*").
			From(table, "").
			Where("year = ?", year).
			OrderBy("rent_number DESC").
			Limit(limit).
			Exec()
	} else {
		err = builder.Select("*").
			From(table, "").
			OrderBy("rent_number DESC").
			Limit(limit).
			Exec()
	}

	if err != nil {
		glogger.Log("Movie.MostRentedList", err.Error())
		return nil, httperror.ErrInternalServerError
	}

	return movies, nil
}

// Get the most rented movie with the requested year.
// year parameter can be empty, in this case it will returns
// the most rented movie all of the times.
func (r *Repository) MostRented(year string) (model.Movie, error) {
	var movie model.Movie
	var err error
	builder := database.NewQueryBuilder(r.db, &movie)

	if year != "" {
		err = builder.Select("*").
			From(table, "").
			Where("year = ?", year).
			OrderBy("rent_number DESC").
			ExecOne()
	} else {
		err = builder.Select("*").
			From(table, "").
			OrderBy("rent_number DESC").
			ExecOne()
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return movie, httperror.ErrNotFound
		}
		glogger.Log("Movie.MostRented", err.Error())
		return movie, httperror.ErrInternalServerError
	}

	return movie, nil
}

// Get the author with the higher rented number.
func (r *Repository) FindBestAuthor() (string, error) {
	var movie model.Movie
	builder := database.NewQueryBuilder(r.db, &movie)

	err := builder.Select("author").
		From(table, "").
		OrderBy("rent_number DESC").
		ExecOne()
	if err != nil {
		if err == sql.ErrNoRows {
			return "", httperror.ErrNotFound
		}
		glogger.Log("Movie.FindBestAuthor", err.Error())
		return "", httperror.ErrInternalServerError
	}

	return movie.Author, nil
}

// Search movies by %title%.
func (r *Repository) FindByTitle(title string) ([]model.Movie, error) {
	movies := []model.Movie{}
	builder := database.NewQueryBuilder(r.db, &movies)

	err := builder.Select("*").
		From(table, "").
		Where("title LIKE ?", "%"+title+"%").
		Exec()
	if err != nil {
		glogger.Log("Movie.FindByTitle", err.Error())
		return nil, httperror.ErrInternalServerError
	}

	return movies, nil
}

// Create a new movie.
func (r *Repository) AddMovie(movie model.Movie) (int64, error) {
	q := "INSERT INTO " + table +
		" (year, rent_number, title, author, editor, `index`, bib, ref, cat_1, cat_2) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := r.db.Exec(q, movie.Year, movie.RentNumber, movie.Title, movie.Author,
		movie.Editor, movie.Index, movie.Bib, movie.Ref, movie.Cat1, movie.Cat2)
	if err != nil {
		glogger.Log("Movie.AddMovie", err.Error())
		return 0, httperror.ErrInternalServerError
	}

	id, err := result.LastInsertId()
	if err != nil {
		glogger.Log("Movie.AddMovie", err.Error())
		return 0, httperror.ErrInternalServerError
	}

	return id, nil
}

// Increment the rented number with the requested title and year.
func (r *Repository) IncrementRentedNumber(title string, year string) error {
	var movie model.Movie
	builder := database.NewQueryBuilder(r.db, &movie)

	// Check if the movie exists
	err := builder.Select("*").From(table, "").
		Where("title = ? AND year = ?", title, year).ExecOne()
	if err != nil {
		if err == sql.ErrNoRows {
			return httperror.ErrNotFound
		}
		glogger.Log("Movie.IncrementRentedNumber", err.Error())
		return httperror.ErrInternalServerError
	}

	q := "UPDATE " + table + " SET rent_number = rent_number + 1 WHERE title=? AND year=?"
	_, err = r.db.Exec(q, title, year)
	if err != nil {
		glogger.Log("Movie.IncrementRentedNumber", err.Error())
		return httperror.ErrInternalServerError
	}

	return nil
}
