package database

import (
	"github.com/hibbannn/grpc-http-boilerplate/app/core/domain"
	"gorm.io/gorm"
)

type Database struct {
	Config domain.DatabaseConfig
	DB     *gorm.DB
}

func NewDatabase(cfg domain.DatabaseConfig) (*Database, error) {
	db := &Database{Config: cfg}
	if err := db.initDatabase(); err != nil {
		return nil, err
	}
	return db, nil
}
