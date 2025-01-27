package mail

import (
	"gopkg.in/mail.v2"
)

type EmailService struct {
	dialer      *mail.Dialer
	fromAddress string
	fromName    string
}

func NewEmailService(host string, port int, username, password, fromAddress, fromName string) *EmailService {
	dialer := mail.NewDialer(host, port, username, password)
	dialer.SSL = false

	return &EmailService{
		dialer:      dialer,
		fromAddress: fromAddress,
		fromName:    fromName,
	}
}

func (s *EmailService) Send(to, subject, body string) error {
	m := mail.NewMessage()
	m.SetAddressHeader("From", s.fromAddress, s.fromName)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	return s.dialer.DialAndSend(m)
}
