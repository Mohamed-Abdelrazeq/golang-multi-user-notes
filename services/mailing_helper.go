package services

import (
	"log"
	"net/smtp"
	"os"
)

func SendWelcomeEmail(email string) {
	// Configuration
	sender := os.Getenv("EMAIL")
	senderAppPassword := os.Getenv("EMAIL_PASSWORD")
	receiver := []string{email}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message
	message := []byte(
		"To: " + receiver[0] + "\r\n" +
			"Subject: Welcome to the new era\r\n" +
			"\r\n" +
			"Welcome to my humble app, I hope you enjoy it and find it useful!",
	)

	// Create authentication
	auth := smtp.PlainAuth("", sender, senderAppPassword, smtpHost)

	// Send actual message
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, sender, receiver, message)
	if err != nil {
		log.Fatal(err)
	}
}
