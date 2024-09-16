package main

import (
	"go-backend/pkg/server"
	"log"
)

func main() {
	srv := server.New()
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
