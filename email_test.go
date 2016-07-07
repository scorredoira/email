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
