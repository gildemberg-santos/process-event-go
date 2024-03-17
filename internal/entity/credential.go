package entity

import (
	"log"
	"os"
	"time"

	"github.com/gildemberg-santos/process-event-go/internal/utils"
)

type Credential struct {
	ClientID string
	SecretID string
}

func NewCredential() *Credential {

	var baseKey string
	if os.Getenv("SECRET_KEY") != "" {
		baseKey = os.Getenv("SECRET_KEY")
	} else {
		baseKey = time.Now().String()
	}
	client_id := utils.ParseStringToSha1(baseKey)
	secret_id := utils.ParseStringToSha512(baseKey)

	log.Println("Client ID: ", client_id)
	log.Println("Secret ID: ", secret_id)

	credential := Credential{
		ClientID: client_id,
		SecretID: secret_id,
	}

	return &credential
}
