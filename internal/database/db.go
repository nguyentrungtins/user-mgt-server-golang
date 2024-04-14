package database

import (
	"user-mgt/internal/app/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(connectionString string, maxOpenConns, maxIdleConns int) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	sqlDB, err := db.DB()
	// Set connection pool configurations
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetMaxIdleConns(maxIdleConns)

	// AutoMigrate schema
	db.AutoMigrate(&entity.User{})

	return db, nil
}
