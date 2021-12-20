package main

import (
	go_api "github.com/AlexKryvets/go-api"
	"github.com/AlexKryvets/go-api/pkg/handler"
	"github.com/AlexKryvets/go-api/pkg/repository"
	"github.com/AlexKryvets/go-api/pkg/service"
	"log"
)

func main() {
	repository := repository.NewRepository()
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	srv := new(go_api.Server)
	if err := srv.Run("8000", handler.InitRouters()); err != nil {
		log.Fatalf("Error")
	}
}
