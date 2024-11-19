package main

import (
	stdMail "net/mail"

	"github.com/trinhminhtriet/mailcli/mail"
	"gopkg.in/jordan-wright/email.v2"
)

func sendMail(recipient Recipient, emailField string, mailer *mail.Mailer,
	debug bool, success chan *email.Email, fail chan error) {

	parsedSender, err := stdMail.ParseAddress(sender)
	if err != nil {
		fail <- err
		return
	}

	parsedTo, err := stdMail.ParseAddress(recipient[emailField])
	if err != nil {
		fail <- err
		return
	}

	message, err := mail.NewMessage(
		parsedSender,
		parsedTo,
		subject,
		files,
		textTemplatePath,
		htmlTemplatePath,
		recipient,
	)
	if err != nil {
		fail <- err
		return
	}

	if !debug {
		if err := mailer.Send(message); err != nil {
			fail <- err
			return
		}
	}

	success <- message
}
