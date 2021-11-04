package models

import (
	"database/sql"
	"fmt"
)

type Movie struct {
	ID         uint32 `json:"id"`
	Year       uint16 `json:"year"`
	RentNumber uint32 `json:"rent_number"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	Editor     string `json:"editor"`
	Index      string `json:"index"`
	Bib        string `json:"bib"`
	Ref        string `json:"ref"`
	Cat1       string `json:"cat_1"`
	Cat2       string `json:"cat_2"`
}

func GetMovies(db *sql.DB) ([]Movie, error) {
	movies := []Movie{}

	rows, err := db.Query("SELECT * FROM movie")
	if err != nil {
		return nil, fmt.Errorf("GetMovies: %v", err)
	}

	defer rows.Close()
	for rows.Next() {
		var movie Movie
		if err := rows.Scan(&movie.ID, &movie.Year, &movie.RentNumber,
			&movie.Title, &movie.Author, &movie.Editor, &movie.Index,
			&movie.Bib, &movie.Ref, &movie.Cat1, &movie.Cat2); err != nil {
			return nil, fmt.Errorf("GetMovies: %v", err)
		}
		movies = append(movies, movie)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetMovies: %v", err)
	}

	return movies, nil
}
