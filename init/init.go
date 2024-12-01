package init

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../configs/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}