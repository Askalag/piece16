package model

type TaskItem struct {
	Id        int
	Title     string
	TreeLevel byte
	ParentId  int

	TimeItems       []TimeItem
	TimeCostTotal   float32
	TimeCostAverage float32
}
