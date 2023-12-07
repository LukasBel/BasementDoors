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
	var subject, body string
	p, s := GetData()

	switch {

	case p >= 50 && s == 0: //raining but not snowing
		percentage := strconv.Itoa(p)
		subject = fmt.Sprintf("%v%% chance of rain. ", percentage)
		body = "Close Basement Door"
	case p >= 50 && s > 0: //snowing
		subject = "ITS SNOWING!"
	default:
		subject = "Something went wrong..."
	}

	message := []byte(subject + body)
	return message
}
