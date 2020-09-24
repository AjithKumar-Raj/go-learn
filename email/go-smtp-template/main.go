package main

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"
)

func main() {

	// Please get fromadd, password and toadd from config value

	// var fromadd string
	// var password string
	// var toadd []string

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", fromAdd, password, smtpHost)

	t, _ := template.ParseFiles("invite.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		UserName string
		Email    string
		Password string
	}{
		UserName: "Ajith",
		Email:    "ajith@gmail.com",
		Password: "WelcomeYou&Me",
	})

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, fromAdd, toAdd, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}
