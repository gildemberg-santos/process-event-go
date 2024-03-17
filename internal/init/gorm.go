package init

import (
	"log"

	"github.com/gildemberg-santos/process-event-go/internal/model"
)

func runGorm() {
	log.Println("Running Gorm")
	db := model.ConectionDB{}
	db.Open()
	db.Migrate(model.Credential{})
	db.Migrate(model.Events{})
	db.Close()
}
