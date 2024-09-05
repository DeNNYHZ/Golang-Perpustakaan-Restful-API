package configs

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgreSQLConn establishes a connection to the PostgreSQL database.
func PostgreSQLConn() *gorm.DB {
	LoadConfig()

	configDB := map[string]string{
		"DB_Username": Config("DB_USERNAME"),
		"DB_Password": Config("DB_PASSWORD"),
		"DB_Port":     Config("DB_PORT"),
		"DB_Host":     Config("DB_ADDRESS"),
		"DB_Name":     Config("DB_NAME"),
	}
	fmt.Println(configDB)
	fmt.Println("Server running well...")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configDB["DB_Host"],
		configDB["DB_Port"],
		configDB["DB_Username"],
		configDB["DB_Password"],
		configDB["DB_Name"])

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
