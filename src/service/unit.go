package service

// UOW Unit of Work...
type UOW struct {
	S1 *Service
}

func NewUOW(s1 *Service) *UOW { // toDo
	return &UOW{
		S1: s1,
	}
}
