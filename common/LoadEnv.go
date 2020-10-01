package common

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
)


func LoadEnv() {
	re := regexp.MustCompile(`^(.*co2)`)
	cwd, _ := os.Getwd()

	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		log.Fatal("Problem loading .env file")
	}
}
