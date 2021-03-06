package model

type TimeItem struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	TimeCost    float32 `db:"time_cost" json:"time_cost,string"`
	TreeLevel   byte    `db:"tree_level" json:"tree_level"`
	ParentId    int     `db:"parent_id" json:"parent_id" binding:"required"`
}

func (m *TimeItem) Valid() bool {
	return true
}
