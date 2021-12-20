package main

import (
	go_api "github.com/AlexKryvets/go-api"
	"github.com/AlexKryvets/go-api/pkg/handler"
	"github.com/AlexKryvets/go-api/pkg/repository"
	"github.com/AlexKryvets/go-api/pkg/service"
	"log"
)

func main() {
	_repository := repository.NewRepository()
	_service := service.NewService(_repository)
	_handler := handler.NewHandler(_service)

	srv := new(go_api.Server)
	if err := srv.Run("8000", _handler.InitRouters()); err != nil {
		log.Fatalf("Error")
	}
}
