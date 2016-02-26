package email

import (
	"github.com/stretchr/testify/assert"
	"net/smtp"
	"strings"
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

func TestAttachment(t *testing.T) {
	m := NewMessage("Hi", "this is the body")
	m.From = "to@example.com"
	m.To = []string{"to@example.com"}
	m.Cc = []string{"to@example.com", "to@example.com"}
	m.Bcc = []string{"to@example.com", "to@example.com"}
	err := m.AttachBuffer("email_test.ics", []byte("test"), false)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, strings.Contains(string(m.Bytes()), "text/calendar"), true,
		"Email message contains calendar",
	)
}
