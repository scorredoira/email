// Copyright 2012 Santiago Corredoira
// Distributed under a BSD-like license.
package email

import (
	"errors"
	"net/smtp"
)

type insecureAuth struct {
	username, password string
}

// InsecureAuth returns an Auth that implements the PLAIN authentication
// mechanism as defined in RFC 4616.
// The returned Auth uses the given username and password to authenticate
// without checking a TLS connection or host like smtp.PlainAuth does.
func InsecureAuth(username, password string) smtp.Auth {
	return &insecureAuth{username, password}
}

func (a *insecureAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	resp := []byte("\x00" + a.username + "\x00" + a.password)
	return "PLAIN", resp, nil
}

func (a *insecureAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		// We've already sent everything.
		return nil, errors.New("unexpected server challenge")
	}
	return nil, nil
}
