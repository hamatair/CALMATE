package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(".env.example"); err != nil {
		log.Fatalf("error load env : %v", err)
	}
}
