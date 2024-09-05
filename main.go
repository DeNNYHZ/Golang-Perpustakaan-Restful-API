package main

import (
	"log"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/api"
	userController "github.com/afrizal423/Golang-Perpustakaan-Restful-API/api/v1/user"
	userService "github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/business/user"
	userRepository "github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/repository/user"

	bukuController "github.com/afrizal423/Golang-Perpustakaan-Restful-API/api/v1/buku"
	bukuService "github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/business/buku"
	bukuRepository "github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/repository/buku"

	_ "github.com/afrizal423/Golang-Perpustakaan-Restful-API/configs"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// PostgreSQL connection string
	dsn := "host=localhost user=postgres password=admin dbname=perpustakaan port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	// Open connection to PostgreSQL using GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Initialize repositories, services, and controllers
	userRepo := userRepository.NewUserRepository(db)
	userServices := userService.NewUserService(userRepo)
	userCon := userController.NewUserController(userServices)

	bukuRepo := bukuRepository.NewBukuRepository(db)
	bukuServices := bukuService.NewBukuService(bukuRepo)
	bukuCon := bukuController.NewBukuController(bukuServices)

	// Setup Fiber app with default configuration
	app := fiber.New()

	// Register API routes
	api.RegisterPath(app, userCon, bukuCon)

	// Start the server
	err = app.Listen(":8000")
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
