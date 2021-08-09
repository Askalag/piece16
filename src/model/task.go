package model

type Task struct {
	Id              int        `db:"id" json:"id"`
	Title           string     `db:"title" json:"title"`
	TreeLevel       byte       `db:"tree_level" json:"tree_level"`
	TreeId          int        `db:"tree_id" json:"tree_id"`
	TaskItems       []TaskItem `json:"task_items"`
	TimeCostTotal   float32    `json:"time_cost_total"`
	TimeCostAverage float32    `json:"time_cost_average"`
}
