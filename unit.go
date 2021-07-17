package piece16

import (
	"github.com/Askalag/piece16/src/service"
)

// UOW Unit of Work...
type UOW struct {
	S1 *service.Service
}

func NewUOW(s1 *service.Service) *UOW {
	return &UOW{
		S1: s1,
	}
}
