package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "server: ", log.LstdFlags)
	server := server.NewServer(logger)
	err := server.Server.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
