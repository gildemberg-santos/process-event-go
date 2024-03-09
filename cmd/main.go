package main

import (
	"github.com/gildemberg-santos/process-event-go/internal/router"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	router.NewRoute().Run()
}
