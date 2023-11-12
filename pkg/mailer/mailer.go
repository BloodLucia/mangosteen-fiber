package mailer

import (
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/kalougata/bookkeeping/pkg/config"
	"net/smtp"
)

type MailOptions struct {
	To      string
	Subject string
	Text    string
}

type Mailer struct {
	conf *config.Config
}

func (m *Mailer) Send(options *MailOptions) error {
	e := email.NewEmail()
	e.From = m.conf.Mailer.From
	e.To = []string{options.To}
	e.Subject = options.Subject
	e.Text = []byte(options.Text)
	auth := smtp.PlainAuth(
		"",
		m.conf.Mailer.UserName,
		m.conf.Mailer.PassWord,
		m.conf.Mailer.Host,
	)

	if err := e.Send(fmt.Sprintf("%s:%d", m.conf.Mailer.Host, m.conf.Mailer.Port), auth); err != nil {
		return err
	}

	return nil
}

func NewMailer(conf *config.Config) *Mailer {
	return &Mailer{conf}
}
