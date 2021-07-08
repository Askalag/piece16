package model

type TaskTree interface {

}

type Task struct {
	id int64
	Title string
	TreeLevel byte

	TaskItems []TaskItem
	TimeCostTotal float32
	TimeCostAverage float32

}