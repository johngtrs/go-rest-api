package model

type Movie struct {
	ID         string `json:"id"`
	Year       uint16 `json:"year" binding:"required"`
	RentNumber uint32 `json:"rent_number" db:"rent_number" binding:"required"`
	Title      string `json:"title" binding:"required"`
	Author     string `json:"author" binding:"required"`
	Editor     string `json:"editor" binding:"required"`
	Index      string `json:"index" binding:"required"`
	Bib        string `json:"bib" binding:"required"`
	Ref        string `json:"ref" binding:"required"`
	Cat1       string `json:"cat_1" db:"cat_1" binding:"required"`
	Cat2       string `json:"cat_2" db:"cat_2" binding:"required"`
}
