package mail

import (
	"go/email-verification/configs"

	"gopkg.in/mail.v2"
)

type EmailService struct {
	dialer      *mail.Dialer
	fromAddress string
	fromName    string
}

func NewEmailService(conf *configs.Config) *EmailService {
	dialer := mail.NewDialer(conf.Mail.Host, conf.Mail.Port, conf.Mail.Username, conf.Mail.Password)
	dialer.SSL = false

	return &EmailService{
		dialer:      dialer,
		fromAddress: conf.Mail.FromAddress,
		fromName:    conf.Mail.FromName,
	}
}

func (service *EmailService) Send(to, subject, body string) error {
	m := mail.NewMessage()
	m.SetAddressHeader("From", service.fromAddress, service.fromName)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	return service.dialer.DialAndSend(m)
}
