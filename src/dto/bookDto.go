package dto

type CreateBookDto struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Gender string `json:"gender"`
	Year   string `json:"year"`
}