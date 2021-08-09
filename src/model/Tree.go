package model

type Tree struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Tasks []Task `json:"tasks"`
}
