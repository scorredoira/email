An easy way to send emails with attachments in Go

**Install**

```bash
go get github.com/scorredoira/email
```

**Usage**

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

**LICENSE**

The MIT License (MIT)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
