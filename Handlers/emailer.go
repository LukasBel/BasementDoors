package Handlers

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/smtp"
	"os"
	"strconv"
)

type EmailAgent struct {
	from     string
	password string
}

func SendMail(to []string) error {

	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err)
	}
	agent := EmailAgent{os.Getenv("FROM"), os.Getenv("PASSWORD")}
	auth := smtp.PlainAuth("", agent.from, agent.password, "smtp.gmail.com")

	err = smtp.SendMail("smtp.gmail.com:587", auth, agent.from, to, Message())
	if err != nil {
		return err
	}

	return nil
}

func Message() []byte {
	percentage := strconv.Itoa(GetData())

	subject := fmt.Sprintf("%v%% chance of rain. ", percentage)
	body := "Close Basement Door"
	message := []byte(subject + body)
	return message
}
