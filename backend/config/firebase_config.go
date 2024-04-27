package config

import (
	"os"
)

type FirebaseConfig struct {
	CredentialsJson string
}

func NewFirebaseConfig() *FirebaseConfig {
	return &FirebaseConfig{
		CredentialsJson: os.Getenv("FIREBASE_CREDENTIALS_JSON"),
	}
}
