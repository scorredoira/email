# email [![Travis-CI](https://travis-ci.org/scorredoira/email.svg?branch=master)](https://travis-ci.org/scorredoira/email) [![GoDoc](https://godoc.org/github.com/scorredoira/email?status.svg)](http://godoc.org/github.com/scorredoira/email) [![Report card](https://goreportcard.com/badge/github.com/scorredoira/email)](https://goreportcard.com/report/github.com/scorredoira/email)

An easy way to send emails with attachments in Go

# Install

```bash
go get github.com/scorredoira/email
```

# Usage

```go
	// compose the message
	m := email.NewMessage("Hi", "this is the body")
	m.From = mail.Address{Name: "From", Address: "from@example.com"}
    m.To = []string{"to@example.com"}
    m.Cc = []string{"cc1@example.com", "cc2@example.com"}
    m.Bcc = []string{"bcc1@example.com", "bcc2@example.com"}

	// add attachments
	if err := m.Attach("email.go"); err != nil {
		log.Fatal(err)
	}

	// send it
	auth := smtp.PlainAuth("", "from@example.com", "pwd", "smtp.zoho.com")	
	if err := email.Send("smtp.zoho.com:587", auth, m); err != nil {
		log.Fatal(err)
	}
```


