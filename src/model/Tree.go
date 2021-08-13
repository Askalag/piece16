package model

type Tree struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	TreeLevel byte   `db:"tree_level" json:"tree_level"`
	Tasks     []Task `json:"tasks"`
}

func (t *Tree) Valid() bool {
	return true
}
