package migrations

import (
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	// Create the peminjamen table first
	err := db.AutoMigrate(&Peminjaman{})
	if err != nil {
		return err
	}

	// Then create the anggota table
	err = db.AutoMigrate(&Anggota{})
	if err != nil {
		return err
	}

	return nil
}

// Define your models
type Peminjaman struct {
	gorm.Model
	IDPeminjaman string
	// Other fields...
}

type Anggota struct {
	gorm.Model
	IDAnggota string `gorm:"primaryKey"`
	Username  string `gorm:"unique"`
	Password  string
	// Foreign key reference to Peminjaman
	IDPeminjaman string
	Peminjaman   Peminjaman `gorm:"foreignKey:IDPeminjaman"`
	// Other fields...
}
