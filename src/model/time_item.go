package model

type TimeItem struct {
	id int64
	Title string
	Description string
	TimeCost float32
	TreeLevel byte
	ParentId int
}