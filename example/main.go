package main

import (
	"log"
	"net/mail"
	"net/smtp"

	"github.com/email"
)

func main() {
	m := email.NewMessage("Hi", "this is the body")
	m.From = mail.Address{Name: "From", Address: "xx@example.com"}
	m.To = []string{"user@example.com"}

	// add attachments
	if err := m.Attach("hhqr-code.png"); err != nil {
		log.Fatal(err)
	}
	m.AddHeader("X-CUSTOMER-id", "xxxxx")

	auth := smtp.PlainAuth("", "user@excample.com", "XXXXX", "smtpdm.aliyun.com")
	if err := email.SendMailWithTlS("smtpdm.aliyun.com:465", auth, m); err != nil {
		log.Fatal(err)
	}
	return
}
