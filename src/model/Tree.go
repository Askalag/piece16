package model

type Tree struct {
	Id    int
	Title string
	Tasks []Task
}

func (t *Tree) ToJSON() map[string]interface{} {
	if t == nil {
		return nil
	}
	return map[string]interface{}{
		"id":    t.Id,
		"title": t.Title,
		"tasks": t.Tasks,
	}
}
