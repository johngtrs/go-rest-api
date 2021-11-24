package model

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title" binding:"required,max=128"`
	Artist string  `json:"artist" binding:"required,max=255"`
	Price  float32 `json:"price" binding:"required,lt=1000000"`
}
