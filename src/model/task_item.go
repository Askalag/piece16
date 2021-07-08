package model

type TaskItem struct {
	id int64
	Title string
	TreeLevel byte
	ParentId int

	TimeItems []TimeItem
	TimeCostTotal float32
	TimeCostAverage float32
}