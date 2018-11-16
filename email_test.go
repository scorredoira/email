package email

import (
	"net/mail"
	"strings"
	"testing"
)

func TestAttachment(t *testing.T) {
	m := NewMessage("Hi", "this is the body")

	if err := m.AttachBuffer("test.ics", []byte("test"), false); err != nil {
		t.Fatal(err)
	}

	if strings.Contains(string(m.Bytes()), "text/calendar") == false {
		t.Fatal("Issue with mailer")
	}
}

func TestHeaders(t *testing.T) {
	m := NewMessage("Hi", "this is the body")

	m.AddHeader("X-HEADER-KEY", "HEADERVAL")

	if strings.Contains(string(m.Bytes()), "X-HEADER-KEY: HEADERVAL\r\n") == false {
		t.Fatal("Could not find header in message")
	}
}

func TestAddTo(t *testing.T) {
	m := NewMessage("Hi", "this is the body")
	names := []string{"firstName", "secondName"}
	addresses := []string{"firstAddress", "secondAddress"}

	firstAddress := mail.Address{Name: names[0], Address: addresses[0]}
	m.AddTo(firstAddress)
	if m.To[0] != firstAddress.String() {
		t.Fatal("Incorrect first element")
	}

	secondAddress := mail.Address{Name: names[1], Address: addresses[1]}
	m.AddTo(secondAddress)
	if m.To[1] != secondAddress.String() {
		t.Fatal("Incorrect second element")
	}
}

func TestAddCc(t *testing.T) {
	m := NewMessage("Hi", "this is the body")
	names := []string{"firstName", "secondName"}
	addresses := []string{"firstAddress", "secondAddress"}

	firstAddress := mail.Address{Name: names[0], Address: addresses[0]}
	m.AddCc(firstAddress)
	if m.Cc[0] != firstAddress.String() {
		t.Fatal("Incorrect first element")
	}

	secondAddress := mail.Address{Name: names[1], Address: addresses[1]}
	m.AddCc(secondAddress)
	if m.Cc[1] != secondAddress.String() {
		t.Fatal("Incorrect second element")
	}
}

func TestAddBcc(t *testing.T) {
	m := NewMessage("Hi", "this is the body")
	names := []string{"firstName", "secondName"}
	addresses := []string{"firstAddress", "secondAddress"}

	firstAddress := mail.Address{Name: names[0], Address: addresses[0]}
	m.AddBcc(firstAddress)
	if m.Bcc[0] != firstAddress.String() {
		t.Fatal("Incorrect first element")
	}

	secondAddress := mail.Address{Name: names[1], Address: addresses[1]}
	m.AddBcc(secondAddress)
	if m.Bcc[1] != secondAddress.String() {
		t.Fatal("Incorrect second element")
	}
}
