package main

import (
	"github.com/hibbannn/grpc-http-boilerplate/app/service/server"
	"github.com/hibbannn/grpc-http-boilerplate/app/util/config"
	"github.com/hibbannn/grpc-http-boilerplate/app/util/database"
	"time"
)

func main() {
	timeStart := time.Now()

	// Initialize configuration
	configs := config.NewConfig()

	cfg, err := configs.InitConfig()
	if err != nil {
		panic(err)
	}

	// Initialize database
	db, err := database.NewDatabase(cfg.Database)
	if err != nil {
		panic(err)
	}

	// Initialize server
	servers := server.NewServer(cfg, db.DB)

	// Start server
	servers.StartServer(timeStart)
}
