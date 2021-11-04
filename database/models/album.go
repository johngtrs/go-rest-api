package models

import (
	"database/sql"
	"fmt"
)

type Album struct {
	ID     uint32  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func GetAlbums(db *sql.DB) ([]Album, error) {
	albums := []Album{}

	rows, err := db.Query("SELECT * FROM album")
	if err != nil {
		return nil, fmt.Errorf("GetAlbums: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("GetAlbums: %v", err)
		}
		albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAlbums: %v", err)
	}

	return albums, nil
}

// albumsByArtist queries for albums that have the specified artist name.
func GetAlbumsByArtist(db *sql.DB, name string) ([]Album, error) {
	// An albums slice to hold data from returned rows.
	albums := []Album{}

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("GetAlbumsByArtist %q: %v", name, err)
	}

	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("GetAlbumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAlbumsByArtist %q: %v", name, err)
	}

	return albums, nil
}

// albumByID queries for the album with the specified ID.
func GetAlbumByID(db *sql.DB, id uint64) (Album, error) {
	// An album to hold data from the returned row.
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("GetAlbumByID %d: no such album", id)
		}
		return alb, fmt.Errorf("GetAlbumByID %d: %v", id, err)
	}
	return alb, nil
}

// AddAlbum adds the specified album to the database,
// returning the album ID of the new entry
func AddAlbum(db *sql.DB, alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("AddAlbum: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("AddAlbum: %v", err)
	}

	return id, nil
}
