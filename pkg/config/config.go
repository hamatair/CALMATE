package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv will load environment variables from .env file if available, otherwise it will load from system environment variables.
func LoadEnv() {
	// Cek apakah file .env ada
	if _, err := os.Stat(".env"); err == nil {
		// Jika file .env ada, muat dari file tersebut
		if err := godotenv.Load(".env"); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	} else {
		// Jika file .env tidak ditemukan, kita tidak melakukan apa-apa karena variabel lingkungan
		// bisa diatur oleh sistem atau Railway
		log.Println("No .env file found, using system environment variables")
	}

	// Jika menggunakan autoload godotenv, Anda dapat memuat variabel dari file .env
	// hanya jika file ada. Namun, jika variabel sudah ada, ini akan mengabaikan.
}

