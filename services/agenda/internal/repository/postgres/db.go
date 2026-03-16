package postgres

import (
	"fmt"

	"go-challenge-agenda/services/agenda/internal/repository/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Open connects to Postgres using the given DSN (e.g. host=localhost user=agenda password=secret dbname=agenda sslmode=disable).
func Open(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open postgres: %w", err)
	}
	return db, nil
}

// Migrate runs AutoMigrate for all domain models (same as SQLite).
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Doctor{},
		&models.WorkingHours{},
		&models.Patient{},
		&models.Reservation{},
		&models.BlockedSlot{},
	)
}
