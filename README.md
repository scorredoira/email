An easy way to send emails with attachments in Go

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
