package init

import (
	"log"
	"os"
	"time"

	"github.com/gildemberg-santos/process-event-go/internal/model"
)

func InitCredential() *model.Credential {
	var baseKey string
	if os.Getenv("SECRET_KEY") != "" {
		baseKey = os.Getenv("SECRET_KEY")
	} else {
		baseKey = time.Now().String()
	}
	client_id := model.ParseStringToSha1(baseKey)
	secret_id := model.ParseStringToSha512(baseKey)

	log.Println("Client ID: ", client_id)
	log.Println("Secret ID: ", secret_id)

	credential := model.Credential{
		ClientID: client_id,
		SecretID: secret_id,
	}

	return &credential
}
