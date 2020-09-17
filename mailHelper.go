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
	mimev := "1.0"
	ctype := "text/plain; charset=\"utf-8\""

	msg := "From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: " + mimev + "\r\n" +
		"Content-Type: " + ctype + "\n\n" +
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
		log.Println(msg)
		return
	}

	log.Println(msg)
}
