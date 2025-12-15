package connection

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"time"
)

func sendEmail(receiver string, firstName string, lastName string) error {

	from := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")

	host := "smtp.gmail.com"
	toList := []string{receiver}

	htmlBody := emailTemplate(lastName)
	subject := "Thank you for reaching out — Dina Taing"

	auth := smtp.PlainAuth("", from, password, host)

	msg := []byte(fmt.Sprintf(
		"From: Dina Taing <your@email.com>\r\n"+
			"To: %s\r\n"+
			"Subject: %s\r\n"+
			"Date: %s\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: text/html; charset=UTF-8\r\n"+
			"\r\n"+
			"%s",
		toList[0],
		subject,
		time.Now().Format(time.RFC1123Z),
		htmlBody,
	))

	err := smtp.SendMail(host, auth, from, toList, msg)

	if err != nil {
		log.Print("[SMTP] error to send an email")
		return err
	}

	return nil

}
