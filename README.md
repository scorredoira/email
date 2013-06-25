An easy way to send emails with attachments

**Install**

```bash
go get github.com/scorredoira/email
```

**Basic Usage**

```go
package main

import (
    "github.com/scorredoira/email"
    "net/smtp"
)

func main() {
    m := email.NewMessage("Hi", "this is the body")
    m.From = "from@example.com"
    m.To = []string{"to@example.com"}
    m.Cc = []string{"cc1@example.com", "cc2@example.com"}
    m.Bcc = []string{"bcc1@example.com", "bcc2@example.com"}

    err = email.Send("smtp.gmail.com:587", smtp.PlainAuth("", "user", "password", "smtp.gmail.com"), m)
}
```

**Send attachments**

```go
m := email.NewMessage("Hi", "this is the body")
m.From = "from@example.com"
m.To = []string{"to@example.com"}
err := m.Attach("picture.png")
if err != nil {
    log.Println(err)
}

err = email.Send("smtp.gmail.com:587", smtp.PlainAuth("", "user", "password", "smtp.gmail.com"), m)
```


**Send unencrypted password**

```go
err = email.SendUnencrypted("mail.example.com:25", "user", "password", m)
```
	
