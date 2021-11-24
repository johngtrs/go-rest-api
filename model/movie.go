package model

type Movie struct {
	ID         string `json:"id"`
	Year       uint16 `json:"year" binding:"required,lt=3000"`
	RentNumber uint32 `json:"rent_number" db:"rent_number" binding:"required,lt=4294967290"`
	Title      string `json:"title" binding:"required,max=255"`
	Author     string `json:"author" binding:"required,max=80"`
	Editor     string `json:"editor" binding:"required,max=125"`
	Index      string `json:"index" binding:"required,max=125"`
	Bib        string `json:"bib" binding:"required,max=20"`
	Ref        string `json:"ref" binding:"required,max=20"`
	Cat1       string `json:"cat_1" db:"cat_1" binding:"required,max=20"`
	Cat2       string `json:"cat_2" db:"cat_2" binding:"required,max=10"`
}
