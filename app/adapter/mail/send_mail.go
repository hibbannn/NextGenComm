package mail

import (
	"crypto/tls"
	"github.com/hibbannn/grpc-http-boilerplate/app/util/constants"
	"gopkg.in/mail.v2"
	"time"
)

// Send mengirim email.
func (m *Mail) Send(to []string, subject, body string) error {
	msg := mail.NewMessage()

	// Set E-Mail sender
	msg.SetHeader(constants.EmailFrom, m.Config.Username)

	// Set E-Mail receivers
	msg.SetHeader(constants.EmailTo, to...)

	// Set E-Mail subject
	msg.SetHeader(constants.EmailSubjectPrefix, subject)

	// Set E-Mail body
	msg.SetBody(constants.TextPlain, body)

	// Settings for SMTP server
	d := mail.NewDialer(m.Config.Server, m.Config.Port, m.Config.Username, m.Config.Password)

	// Optional: Configure TLS/SSL
	if m.Config.SSL {
		d.SSL = true
	} else if m.Config.TLS {
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}

	// Optional: Configure timeout, retries, etc.
	d.Timeout = time.Duration(m.Config.Timeout) * time.Second
	d.RetryFailure = m.Config.MaxRetries > 0

	// Now send E-Mail
	return d.DialAndSend(msg)
}
