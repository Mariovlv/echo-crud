package models

type Album struct {
	ID 	   string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Date   string `json:"date"`
}
