package database

import (
	"fmt"
	"github.com/hibbannn/grpc-http-boilerplate/app/util/constants"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

func (d *Database) initDatabase() error {
	dsn := d.buildDSN()
	db, err := gorm.Open(d.determineDialect(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	d.DB = db
	return d.configureDBPool()
}

func (d *Database) buildDSN() string {

	switch d.Config.Driver {
	case constants.Postgres:
		return fmt.Sprintf(constants.PostgresDSN, d.Config.Host, d.Config.User, d.Config.Password, d.Config.Name, d.Config.Port, d.Config.SSLMode, d.Config.Schema)
	case constants.Mysql:
		return fmt.Sprintf(constants.MysqlDSN, d.Config.User, d.Config.Password, d.Config.Host, d.Config.Port, d.Config.Name)
	case constants.Sqlite:
		return fmt.Sprintf(constants.SqliteDSN, d.Config.Name)
	default:
		panic("unsupported database driver") // Consider handling this more gracefully
	}
}

func (d *Database) determineDialect(dsn string) gorm.Dialector {
	switch d.Config.Driver {
	case constants.Postgres:
		return postgres.Open(dsn)
	case constants.Mysql, constants.Mariadb:
		return mysql.Open(dsn)
	case constants.Sqlite:
		return sqlite.Open(dsn)
	default:
		panic("unsupported database driver")
	}
}

func (d *Database) configureDBPool() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(d.Config.MaxIdle)
	sqlDB.SetMaxOpenConns(d.Config.MaxOpen)
	sqlDB.SetConnMaxLifetime(time.Duration(d.Config.MaxLife) * time.Minute)
	sqlDB.SetConnMaxIdleTime(time.Duration(d.Config.MaxIdleLife) * time.Minute)

	return nil
}
