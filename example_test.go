package email_test

import (
	"log"
	"net/mail"
	"net/smtp"

	"github.com/scorredoira/email"
)

func Example() {
	// compose the message
	m := email.NewMessage("Hi", "this is the body")
	m.From = mail.Address{Name: "From", Address: "from@example.com"}
	m.To = []string{"to@example.com"}

	// add attachments
	if err := m.Attach("email.go"); err != nil {
		log.Fatal(err)
	}

	// send it
	auth := smtp.PlainAuth("", "from@example.com", "pwd", "smtp.zoho.com")
	if err := email.Send("smtp.zoho.com:587", auth, m); err != nil {
		log.Fatal(err)
	}
}
