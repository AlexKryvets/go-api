package service

import "github.com/AlexKryvets/go-api/pkg/repository"

type Authorization interface {
}

type Service struct {
	Authorization
}

func NewService(repository *repository.Repository) *Service {
	return &Service{}
}
