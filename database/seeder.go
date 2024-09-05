package seed

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"default:'user'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	dsn := "host=localhost user=postgres password=admin dbname=perpustakaan port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// AutoMigrate the User model
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Seed admin user
	admin := User{
		Username: "admin",
		Password: "adminpassword", // Ensure to hash the password in a real application
		Role:     "admin",
	}

	// Check if the user already exists
	var existingUser User
	result := db.Where("username = ?", admin.Username).First(&existingUser)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// User does not exist, create it
			if err := db.Create(&admin).Error; err != nil {
				log.Fatalf("failed to seed admin user: %v", err)
			}
			log.Println("Admin user seeded successfully")
		} else {
			log.Fatalf("failed to check for existing user: %v", result.Error)
		}
	} else {
		log.Println("Admin user already exists")
	}
}
