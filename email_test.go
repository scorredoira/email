package email

import (
	"net/smtp"
	"testing"
)

func TestSend(t *testing.T) {
	m := NewMessage("Hi", "this is the body")
	m.From = "to@example.com"
	m.To = []string{"to@example.com"}
	m.Cc = []string{"to@example.com", "to@example.com"}
	m.Bcc = []string{"to@example.com", "to@example.com"}

	err := m.Attach("email_test.go")
	if err != nil {
		panic(err)
	}

	err = Send("smtp.gmail.com:587", smtp.PlainAuth("", "user", "passoword", "smtp.gmail.com"), m)
	if err != nil {
		panic(err)
	}
}
