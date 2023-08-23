package helpers

import (
	"log"
	"net/smtp"
	"os"
)

func SendVerificationMail(email string) {

	// Configuration
	sender := os.Getenv("EMAIL")
	senderAppPassword := os.Getenv("EMAIL_PASSWORD")
	receiver := []string{email}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("To: " + receiver[0] + "\r\n" +
		"Subject: Notes App Verification Code\r\n" +
		"\r\n" +
		"THIS IS YOUR VERIFICATION CODE\r\n")

	// Create authentication
	auth := smtp.PlainAuth("", sender, senderAppPassword, smtpHost)

	// Send actual message
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, sender, receiver, message)
	if err != nil {
		log.Fatal(err)
	}
}
