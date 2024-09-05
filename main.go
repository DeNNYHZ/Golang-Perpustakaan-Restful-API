package main

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/api"
	userController "github.com/afrizal423/Golang-Perpustakaan-Restful-API/api/v1/user"
	userService "github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/business/user"
	userRepository "github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/repository/user"
	"log"

	bukuController "github.com/afrizal423/Golang-Perpustakaan-Restful-API/api/v1/buku"
	bukuService "github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/business/buku"
	bukuRepository "github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/repository/buku"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/migrations"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=admin dbname=perpustakaan port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Run migrations
	if err := migrations.RunMigrations(db); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	userRepo := userRepository.NewUserRepository(db)
	userServices := userService.NewUserService(userRepo)
	userCon := userController.NewUserController(userServices)

	bukuRepo := bukuRepository.NewBukuRepository(db)
	bukuServices := bukuService.NewBukuService(bukuRepo)
	bukuCon := bukuController.NewBukuController(bukuServices)

	app := fiber.New()

	// Define root route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the API!")
	})

	api.RegisterPath(app, userCon, bukuCon)

	err = app.Listen(":8000")
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
