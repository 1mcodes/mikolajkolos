package sender

import (
	"fmt"
	"net/smtp"
)

type Mailer struct {
	from, address string
	auth          smtp.Auth
}

func NewMailer(address, port, login, password string) Mailer {
	return Mailer{
		from:    login,
		address: address + ":" + port,
		auth:    smtp.PlainAuth("login", login, password, address),
	}
}

func (m Mailer) Send(to []string, msg []byte) error {
	return smtp.SendMail(m.address, m.auth, m.from, to, msg)
}

func (m Mailer) PrepareMessage(who, whom, tip string) string {
	return fmt.Sprintf(MAIL_TEMPLATE, who, who, whom, tip)
}
