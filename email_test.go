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

	m.AddTo(mail.Address{Name: names[0], Address: addresses[0]})
	if m.To[0] != names[0]+" "+"<"+addresses[0]+">" {
		t.Fatal("Incorrect first element")
	}

	m.AddTo(mail.Address{Name: names[1], Address: addresses[1]})
	if m.To[1] != names[1]+" "+"<"+addresses[1]+">" {
		t.Fatal("Incorrect second element")
	}
}

func TestAddCc(t *testing.T) {
	m := NewMessage("Hi", "this is the body")
	names := []string{"firstName", "secondName"}
	addresses := []string{"firstAddress", "secondAddress"}

	m.AddCc(mail.Address{Name: names[0], Address: addresses[0]})
	if m.Cc[0] != names[0]+" "+"<"+addresses[0]+">" {
		t.Fatal("Incorrect first element")
	}

	m.AddCc(mail.Address{Name: names[1], Address: addresses[1]})
	if m.Cc[1] != names[1]+" "+"<"+addresses[1]+">" {
		t.Fatal("Incorrect second element")
	}
}

func TestAddBcc(t *testing.T) {
	m := NewMessage("Hi", "this is the body")
	names := []string{"firstName", "secondName"}
	addresses := []string{"firstAddress", "secondAddress"}

	m.AddBcc(mail.Address{Name: names[0], Address: addresses[0]})
	if m.Bcc[0] != names[0]+" "+"<"+addresses[0]+">" {
		t.Fatal("Incorrect first element")
	}

	m.AddBcc(mail.Address{Name: names[1], Address: addresses[1]})
	if m.Bcc[1] != names[1]+" "+"<"+addresses[1]+">" {
		t.Fatal("Incorrect second element")
	}
}
