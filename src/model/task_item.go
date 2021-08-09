package model

type TaskItem struct {
	Id              int        `json:"id"`
	Title           string     `json:"title"`
	TreeLevel       byte       `db:"tree_level" json:"tree_level"`
	ParentId        int        `db:"parent_id" json:"parent_id"`
	TimeItems       []TimeItem `json:"time_items"`
	TimeCostTotal   float32    `json:"time_cost_total"`
	TimeCostAverage float32    `json:"time_cost_average"`
}
