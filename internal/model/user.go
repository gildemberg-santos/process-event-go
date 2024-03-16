package model

import (
	"log"
	"os"
	"time"
)

type Credential struct {
	// gorm.Model
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
	client_id := ParseStringToSha1(baseKey)
	secret_id := ParseStringToSha512(baseKey)

	log.Println("Client ID: ", client_id)
	log.Println("Secret ID: ", secret_id)

	credential := Credential{
		ClientID: client_id,
		SecretID: secret_id,
	}

	return &credential
}
