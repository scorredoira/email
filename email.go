// Copyright 2012 Santiago Corredoira
// Distributed under a BSD-like license.
package email

import (
	"bytes"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"net/smtp"
	"path/filepath"
	"strings"
)

type Message struct {
	From        string
	To          []string
	Cc          []string
	Bcc         []string
	Subject     string
	Body        string
	Attachments map[string][]byte
}

func (m *Message) Attach(file string) error {
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

func (m *Message) Tolist() []string {
	tolist := m.To
	for _, cc := range m.Cc {
		tolist = append(tolist, cc)
	}
	for _, bcc := range m.Bcc {
		tolist = append(tolist, bcc)
	}
	return tolist
}

func (m *Message) Bytes() []byte {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("From: " + m.From + "\n")
	buf.WriteString("To: " + strings.Join(m.To, ",") + "\n")
	if len(m.Cc) > 0 {
		buf.WriteString("Cc: " + strings.Join(m.Cc, ",") + "\n")
	}
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

func Send(addr string, auth smtp.Auth, m *Message) error {
	return smtp.SendMail(addr, auth, m.From, m.Tolist(), m.Bytes())
}

func SendUnencrypted(addr, user, password string, m *Message) error {
	auth := UnEncryptedAuth(user, password)
	return smtp.SendMail(addr, auth, m.From, m.Tolist(), m.Bytes())
}

type unEncryptedAuth struct {
	username, password string
}

// InsecureAuth returns an Auth that implements the PLAIN authentication
// mechanism as defined in RFC 4616.
// The returned Auth uses the given username and password to authenticate
// without checking a TLS connection or host like smtp.PlainAuth does.
func UnEncryptedAuth(username, password string) smtp.Auth {
	return &unEncryptedAuth{username, password}
}

func (a *unEncryptedAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	resp := []byte("\x00" + a.username + "\x00" + a.password)
	return "PLAIN", resp, nil
}

func (a *unEncryptedAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		// We've already sent everything.
		return nil, errors.New("unexpected server challenge")
	}
	return nil, nil
}
