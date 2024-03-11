package server

import (
	"github.com/hibbannn/grpc-http-boilerplate/app/adapter/cache"
	"github.com/hibbannn/grpc-http-boilerplate/app/adapter/client"
	"github.com/hibbannn/grpc-http-boilerplate/app/adapter/mail"
	"github.com/hibbannn/grpc-http-boilerplate/app/adapter/repository"
	"github.com/hibbannn/grpc-http-boilerplate/app/core/contract"
	"github.com/hibbannn/grpc-http-boilerplate/app/core/domain"
	"github.com/hibbannn/grpc-http-boilerplate/app/core/rule"
	"github.com/hibbannn/grpc-http-boilerplate/app/util/logs"
	"github.com/hibbannn/grpc-http-boilerplate/app/util/middlewares"
	"github.com/hibbannn/grpc-http-boilerplate/app/util/redis"
	"github.com/hibbannn/grpc-http-boilerplate/app/util/security"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"net/http"
)

type Server struct {
	grpcServer *grpc.Server
	httpServer *http.Server
	config     *domain.Config
	db         *gorm.DB
	middleware contract.MiddlewareContract
	repo       contract.RepositoryContract
	cache      contract.CacheContract
	mail       contract.MailContract
	client     contract.ClientContract
	rule       contract.RuleContract
	security   contract.SecurityContract
	log        contract.LogContract
}

// NewServer initializes a new server with the provided configurations.
func NewServer(cfg *domain.Config, db *gorm.DB) *Server {

	// Initialize logs
	log := logs.NewLogs(cfg.Logging)

	err := log.InitializeLogs()
	if err != nil {
		panic(err)
	}

	redises := redis.NewRedis(cfg.Cache)

	// Register services
	repos := repository.NewRepository(db)
	caches := cache.NewCache(redises)
	mails := mail.NewMail(cfg.Mail)
	clients := client.NewClient(cfg)
	rules := rule.NewRule(cfg)
	securities := security.NewSecurity(cfg.JWT, cfg.Encryption)
	middleware, err := middlewares.NewMiddleware(cfg.RateLimiter)
	if err != nil {
		panic(err)
	}

	s := &Server{
		config:     cfg,
		db:         db,
		middleware: middleware,
		repo:       repos,
		cache:      caches,
		mail:       mails,
		client:     clients,
		rule:       rules,
		security:   securities,
	}
	s.initializeGRPCServer()
	s.initializeHTTPServer()
	return s
}
