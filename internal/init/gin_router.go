package init

import (
	"log"

	"github.com/gildemberg-santos/process-event-go/internal/router"
)

func runGinRouter() {
	log.Println("Running Gin Router")
	router.NewRoute().Run()
}
