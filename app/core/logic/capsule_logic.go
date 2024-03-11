package logic

import (
	"github.com/hibbannn/grpc-http-boilerplate/app/core/contract"
	"github.com/hibbannn/grpc-http-boilerplate/stub/serviceUser"
)

type Logic struct {
	repository contract.RepositoryContract
	cache      contract.CacheContract
	mail       contract.MailContract
	client     contract.ClientContract
	rule       contract.RuleContract
	security   contract.SecurityContract
	log        contract.LogContract
	serviceUser.UserServiceServer
}

func NewLogic(repository contract.RepositoryContract, cache contract.CacheContract, mail contract.MailContract, client contract.ClientContract, rule contract.RuleContract, security contract.SecurityContract, log contract.LogContract) *Logic {
	return &Logic{
		repository: repository,
		cache:      cache,
		mail:       mail,
		client:     client,
		rule:       rule,
		security:   security,
		log:        log,
	}
}

func (l *Logic) mustEmbedUnimplementedUserServiceServer() {
}
