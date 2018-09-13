package email

import (
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