package movie

type Movie struct {
	ID         string `json:"id"`
	Year       uint16 `json:"year"`
	RentNumber uint32 `json:"rent_number" db:"rent_number"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	Editor     string `json:"editor"`
	Index      string `json:"index"`
	Bib        string `json:"bib"`
	Ref        string `json:"ref"`
	Cat1       string `json:"cat_1" db:"cat_1"`
	Cat2       string `json:"cat_2" db:"cat_2"`
}
