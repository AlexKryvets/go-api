package main

import (
	go_api "github.com/AlexKryvets/go-api"
	handler "github.com/AlexKryvets/go-api/pkg/handler"
	"log"
)

func main() {
	handler := new(handler.Handler)
	srv := new(go_api.Server)
	if err := srv.Run("8000", handler.InitRouters()); err != nil {
		log.Fatalf("Error")
	}
}
