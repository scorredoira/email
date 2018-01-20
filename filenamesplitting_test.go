package email

import (
	"net/mail"
	"strings"
	"testing"
)

func TestFilenameSplitting(t *testing.T) {

	message := NewMessage("Subject", "Text")
	message.From = mail.Address{"From", "from@example.com"}
	message.To = []string{"to@example.com"}

	longstring := "123456789012345678901234567890123456789012345678901234567890123456789012345678901234"
	filecontent := []byte("Hello World")

	message.AttachBuffer(longstring[0:28], filecontent, false) // 28 -> 40 (base64 and utf8-prefix)
	message.AttachBuffer(longstring[0:32], filecontent, false) // 32 -> 44
	message.AttachBuffer(longstring[0:36], filecontent, false) // 36 -> 48
	message.AttachBuffer(longstring[0:40], filecontent, false) // 40 -> 56
	message.AttachBuffer(longstring[0:80], filecontent, false) // 80 -> 108
	message.AttachBuffer(longstring[0:84], filecontent, false) // 84 -> 112

	output := string(message.Bytes())
	t.Log(output)

	if strings.Count(output, "filename=") != 1 {
		t.Fatal("Wrong number of filename=")
	}

	if strings.Count(output, "filename*0=") != 5 {
		t.Fatal("Wrong number of filename*0=")
	}

	if strings.Count(output, "filename*1=") != 3 {
		t.Fatal("Wrong number of filename*1=")
	}
	if strings.Count(output, "filename*2=") != 1 {
		t.Fatal("Wrong number of filename*2=")
	}
}
