package migrations

import (
	"crud-go/models"
	"gorm.io/gorm"
)

func MigrateAll(db *gorm.DB) {
	db.AutoMigrate(
		&models.Doctors{},
		&models.Students{},
		&models.StaffMembers{},
		&models.User{},
		// Add all your models here
	)
}
