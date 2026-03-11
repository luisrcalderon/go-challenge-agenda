package sqlite

import (
	"fmt"

	"go-challenge-agenda/services/agenda/internal/repository/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Open(path string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open sqlite: %w", err)
	}
	return db, nil
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Doctor{},
		&models.WorkingHours{},
		&models.Patient{},
		&models.Reservation{},
		&models.BlockedSlot{},
	)
}
