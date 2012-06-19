// Copyright 2012 Santiago Corredoira
// Distributed under a BSD-like license.
package email

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"net/smtp"
	"path/filepath"
)

type Message struct {
	Subject     string
	Body        string
	Attachments map[string][]byte
}

func (m *Message) AddAttachment(file string) error {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	_, fileName := filepath.Split(file)
	m.Attachments[fileName] = b
	return nil
}

// NewMessage returns a new Message that can compose an email with attachments
func NewMessage(subject string, body string) *Message {
	m := &Message{Subject: subject, Body: body}
	m.Attachments = make(map[string][]byte)
	return m
}

func (m *Message) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	buf.WriteString("Subject: " + m.Subject + "\n")
	buf.WriteString("MIME-Version: 1.0\n")

	boundary := "f46d043c813270fc6b04c2d223da"

	if len(m.Attachments) > 0 {
		buf.WriteString("Content-Type: multipart/mixed; boundary=" + boundary + "\n")
		buf.WriteString("--" + boundary + "\n")
	}

	buf.WriteString("Content-Type: text/plain; charset=utf-8\n")
	buf.WriteString(m.Body)

	if len(m.Attachments) > 0 {
		for k, v := range m.Attachments {
			buf.WriteString("\n\n--" + boundary + "\n")
			buf.WriteString("Content-Type: application/octet-stream\n")
			buf.WriteString("Content-Transfer-Encoding: base64\n")
			buf.WriteString("Content-Disposition: attachment; filename=\"" + k + "\"\n\n")

			b := make([]byte, base64.StdEncoding.EncodedLen(len(v)))
			base64.StdEncoding.Encode(b, v)
			buf.Write(b)
			buf.WriteString("\n--" + boundary)
		}

		buf.WriteString("--")
	}

	return buf.Bytes()
}

func Send(addr string, auth smtp.Auth, from string, to []string, m *Message) error {
	return smtp.SendMail(addr, auth, from, to, m.Bytes())
}
