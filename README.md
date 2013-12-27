An easy way to send emails with attachments

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
