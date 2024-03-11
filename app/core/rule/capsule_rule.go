package rule

import "github.com/hibbannn/grpc-http-boilerplate/app/core/domain"

type Rule struct {
	Config *domain.Config
}

func NewRule(config *domain.Config) *Rule {
	return &Rule{
		Config: config,
	}
}
