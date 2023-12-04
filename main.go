package main

import (
	"example.com/Handlers"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	probability := Handlers.GetData()
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err)
	}

	emails := os.Getenv("TO")
	emailAddresses := strings.Split(emails, ",")

	for {
		if probability >= 50 {
			Handlers.SendMail(emailAddresses)
		}
		time.Sleep(time.Minute * 30)
	}

}
