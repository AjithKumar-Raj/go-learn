package main

import (
	"fmt"
	"net/smtp"
)

func main() {

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587" //587

	// Message.
	message := []byte("This is a test email message.")

	// Authentication.
	auth := smtp.PlainAuth("", fromAdd, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, fromAdd, toAdd, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
