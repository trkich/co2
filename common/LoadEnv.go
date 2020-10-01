package common

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
)


func LoadEnv() {
	re := regexp.MustCompile(`^(.*co2)`)
	cwd, _ := os.Getwd()

	rootPath := re.Find([]byte(cwd))

	if flag.Lookup("test.v") == nil {
		err := godotenv.Load(string(rootPath) + `/.env`)

		if err != nil {
			log.Fatal("Problem loading .env file")
		}

	} else {
		err := godotenv.Load(string(rootPath) + `/.env.test`)

		if err != nil {
			log.Fatal("Problem loading .env.test file")
		}
	}



}
