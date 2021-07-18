package model

type Task struct {
	Id        int64  `db:"id"`
	Title     string `db:"title"`
	TreeLevel byte   `db:"tree_level"`

	TaskItems       []TaskItem
	TimeCostTotal   float32
	TimeCostAverage float32
}

func ToJSONArr(src []Task) []map[string]interface{} {
	if src == nil || len(src) == 0 {
		return nil
	}
	res := make([]map[string]interface{}, len(src))
	for i, v := range src {
		res[i] = v.ToJSON()
	}
	return res
}

func (t *Task) ToJSON() map[string]interface{} {
	if t == nil {
		return nil
	}
	return map[string]interface{}{
		"id":              t.Id,
		"title":           t.Title,
		"treeLevel":       t.TreeLevel,
		"taskItems":       t.TaskItems,
		"timeCostTotal":   t.TimeCostTotal,
		"timeCostAverage": t.TimeCostAverage,
	}
}
