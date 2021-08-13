package validator

import "github.com/Askalag/piece16/src/model"

func ValidId(id *int) bool {
	if id == nil || *id <= 0 {
		return false
	}
	return true
}

func ValidateTask(task model.Task) bool {
	return true
}
