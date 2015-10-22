// Copyright 2012 Santiago Corredoira
// Distributed under a BSD-like license.
package email

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/smtp"
	"path/filepath"
	"strings"
	"time"
)

const lineMaxLen = 76

type AttachmentWriter struct {
	w    io.Writer
	i    int
	line [78]byte
}

// NewAttachmentWriter returns a new AttachmentWriter that writes base64 encoded
// attachments to the underlying writer in lines no more than 76 characters each.
func NewAttachmentWriter(w io.Writer) *AttachmentWriter {
	return &AttachmentWriter{
		w: w,
	}
}

// Write writes the bytes from p to the underlying writer in lines no more than
// 76 characters each.
func (w *AttachmentWriter) Write(p []byte) (n int, err error) {
	for i, b := range p {
		w.line[w.i] = b
		w.i++
		n = i

		if w.i == lineMaxLen {
			w.insertCRLF()
		}
	}

	if err := w.flush(); err != nil {
		return n, err
	}

	return len(p), nil
}

func (w *AttachmentWriter) insertCRLF() error {
	w.line[w.i] = '\r'
	w.line[w.i+1] = '\n'
	w.i += 2

	return w.flush()
}

func (w *AttachmentWriter) flush() error {
	if _, err := w.w.Write(w.line[:w.i]); err != nil {
		return err
	}

	w.i = 0

	return nil
}

type Attachment struct {
	Filename string
	Data     []byte
	Inline   bool
}

type Message struct {
	From            string
	To              []string
	Cc              []string
	Bcc             []string
	Subject         string
	Body            string
	BodyContentType string
	Attachments     map[string]*Attachment
}

func (m *Message) attach(file string, inline bool) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	_, filename := filepath.Split(file)

	m.Attachments[filename] = &Attachment{
		Filename: filename,
		Data:     data,
		Inline:   inline,
	}

	return nil
}

func (m *Message) Attach(file string) error {
	return m.attach(file, false)
}

func (m *Message) Inline(file string) error {
	return m.attach(file, true)
}

func newMessage(subject string, body string, bodyContentType string) *Message {
	m := &Message{Subject: subject, Body: body, BodyContentType: bodyContentType}

	m.Attachments = make(map[string]*Attachment)

	return m
}

// NewMessage returns a new Message that can compose an email with attachments
func NewMessage(subject string, body string) *Message {
	return newMessage(subject, body, "text/plain")
}

// NewMessage returns a new Message that can compose an HTML email with attachments
func NewHTMLMessage(subject string, body string) *Message {
	return newMessage(subject, body, "text/html")
}

// ToList returns all the recipients of the email
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

// Bytes returns the mail data
func (m *Message) Bytes() []byte {
	buf := bytes.NewBuffer(nil)

	buf.WriteString("From: " + m.From + "\n")

	t := time.Now()
	buf.WriteString("Date: " + t.Format(time.RFC822) + "\n")

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

	buf.WriteString(fmt.Sprintf("Content-Type: %s; charset=utf-8\n\n", m.BodyContentType))
	buf.WriteString(m.Body)
	buf.WriteString("\n")

	if len(m.Attachments) > 0 {
		for _, attachment := range m.Attachments {
			buf.WriteString("\n\n--" + boundary + "\n")

			if attachment.Inline {
				buf.WriteString("Content-Type: message/rfc822\n")
				buf.WriteString("Content-Disposition: inline; filename=\"" + attachment.Filename + "\"\n\n")

				buf.Write(attachment.Data)
			} else {
				buf.WriteString("Content-Type: application/octet-stream\n")
				buf.WriteString("Content-Transfer-Encoding: base64\n")
				buf.WriteString("Content-Disposition: attachment; filename=\"" + attachment.Filename + "\"\n\n")

				b := make([]byte, base64.StdEncoding.EncodedLen(len(attachment.Data)))
				base64.StdEncoding.Encode(b, attachment.Data)

				aw := NewAttachmentWriter(buf)
				aw.Write(b)
			}

			buf.WriteString("\n--" + boundary)
		}

		buf.WriteString("--")
	}

	return buf.Bytes()
}

func Send(addr string, auth smtp.Auth, m *Message) error {
	return smtp.SendMail(addr, auth, m.From, m.Tolist(), m.Bytes())
}
