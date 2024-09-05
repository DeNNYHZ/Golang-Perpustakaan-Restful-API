package api

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/api/v1/buku"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/api/v1/user"
	"github.com/gofiber/fiber/v2"
)

// RegisterPath sets up all the routes for the API.
func RegisterPath(f *fiber.App, userCon *user.Controller, bukuCon *buku.Controller) {
	// Group routes under /api/v1
	apiGroup := f.Group("/api/v1")

	// User routes
	userRoutes := apiGroup.Group("/users")
	userRoutes.Post("/login", userCon.Login)

	// Buku routes
	bukuRoutes := apiGroup.Group("/bukus")

	// Jenis Buku routes
	jenisBukuRoutes := bukuRoutes.Group("/jenis")
	jenisBukuRoutes.Get("/", bukuCon.GetAllJenisBuku)
	jenisBukuRoutes.Get("/:id", bukuCon.GetJenisBukuById)
	jenisBukuRoutes.Post("/", bukuCon.CreateJenisBuku) // Fixed: Added leading '/'
	jenisBukuRoutes.Put("/:id", bukuCon.UpdateJenisBuku)
	jenisBukuRoutes.Delete("/:id", bukuCon.DeleteJenisBuku)

	// Penerbit Buku routes
	penerbitBukuRoutes := bukuRoutes.Group("/penerbit")
	penerbitBukuRoutes.Get("/", bukuCon.GetAllPenerbitBuku)
	penerbitBukuRoutes.Get("/:id", bukuCon.GetPenerbitBukuById)
	penerbitBukuRoutes.Post("/", bukuCon.CreatePenerbitBuku) // Fixed: Added leading '/'
	penerbitBukuRoutes.Put("/:id", bukuCon.UpdatePenerbitBuku)
	penerbitBukuRoutes.Delete("/:id", bukuCon.DeletePenerbitBuku)

	// Penulis Buku routes
	penulisBukuRoutes := bukuRoutes.Group("/penulis")
	penulisBukuRoutes.Get("/", bukuCon.GetAllPenulisBuku)
	penulisBukuRoutes.Get("/:id", bukuCon.GetPenulisBukuById)
	penulisBukuRoutes.Post("/", bukuCon.CreatePenulisBuku) // Fixed: Added leading '/'
	penulisBukuRoutes.Put("/:id", bukuCon.UpdatePenulisBuku)
	penulisBukuRoutes.Delete("/:id", bukuCon.DeletePenulisBuku)
}
