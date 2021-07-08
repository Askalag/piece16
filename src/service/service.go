package service

import "github.com/Askalag/piece16/src/repository"

type Service struct {
	Tree *Tree
}

func NewService(r *repository.Repo) *Service {
	return &Service{
		Tree: NewTreeService(r),
	}
}