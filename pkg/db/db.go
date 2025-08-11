package db

import (
	"fmt"
	"time"

	"github.com/stephenZ22/mini_dash/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Build the connection string
func ConnectPostgres(host string, port int, user, password, dbname string, sslmode string, max_connects int) (*gorm.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	// Connect to the PostgreSQL database
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		logger.MiniLogger().Error("Failed to connect to database", "error", err)
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		logger.MiniLogger().Error("Failed to get database instance", "error", err)
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	sqlDb.SetMaxIdleConns(10)               // Set maximum number of idle connections
	sqlDb.SetMaxOpenConns(max_connects)     // Set maximum number of open connections
	sqlDb.SetConnMaxLifetime(1 * time.Hour) // Set maximum connection lifetime (0 means no limit

	return db, nil
}
