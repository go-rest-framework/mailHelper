// Package mailHelper provides mail send with gmail smtp
package mailHelper

import (
	"log"
	"net/smtp"
)

type Mailer struct {
	Email string
	Pass  string
}

func (m *Mailer) Send(to, subject, body string) {
	from := m.Email
	pass := m.Pass

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from,
		[]string{to},
		[]byte(msg),
	)

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
}
