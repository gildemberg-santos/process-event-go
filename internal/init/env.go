package init

import (
	"log"

	"github.com/joho/godotenv"
)

func runEnv() {
	log.Println("Running Variables Environment")
	godotenv.Load()
}
