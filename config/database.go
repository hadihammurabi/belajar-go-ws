package config

import (
	"errors"
	"os"

	"github.com/hadihammurabi/belajar-go-ws/platform/database"

	"gorm.io/gorm"
)

const (
	driverPostgresql = "postgresql"
	driverSqlite     = "sqlite"
)

// ConfigureDatabase func
func ConfigureDatabase() (*gorm.DB, error) {
	driver := os.Getenv("DB_DRIVER")
	if driver == "" {
		driver = driverPostgresql
	}

	if driver == driverPostgresql {
		db, err := database.ConfigurePostgresql()
		return db, err
	} else if driver == driverSqlite {
		db, err := database.ConfigureSqlite()
		return db, err
	}

	return nil, errors.New("unknown database driver")
}
