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
	m.AddTo(mail.Address{Name: "someToName", Address: "to@example.com"})
	m.AddCc(mail.Address{Name: "someCcName", Address: "cc@example.com"})
	m.AddBcc(mail.Address{Name: "someBccName", Address: "bcc@example.com"})

	// add attachments
	if err := m.Attach("email.go"); err != nil {
		log.Fatal(err)
	}

	// add headers
	m.AddHeader("X-CUSTOMER-id", "xxxxx")

	// send it
	auth := smtp.PlainAuth("", "from@example.com", "pwd", "smtp.zoho.com")
	if err := email.Send("smtp.zoho.com:587", auth, m); err != nil {
		log.Fatal(err)
	}
}
