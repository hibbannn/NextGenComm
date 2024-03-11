package mail

import "github.com/hibbannn/grpc-http-boilerplate/app/core/domain"

type Mail struct {
	Config domain.MailConfig
}

func NewMail(config domain.MailConfig) *Mail {
	return &Mail{
		Config: config,
	}
}
