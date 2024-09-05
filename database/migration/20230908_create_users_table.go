package migration

import (
	"gorm.io/gorm"
	"time"
)

// Run migrations for the users table
func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}

// User model for migration
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"default:'user'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
