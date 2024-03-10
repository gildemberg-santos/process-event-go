package model

import "gorm.io/gorm"

type Credential struct {
	gorm.Model
	ClientID string
	SecretID string
}
