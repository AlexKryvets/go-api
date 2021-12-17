package main

import (
	http "github.com/AlexKryvets/go-api"
	"log"
)

func main() {
	srv := new(http.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("Error")
	}
}
