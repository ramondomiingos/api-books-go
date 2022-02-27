package migrations

import (
	"api-books-go/models"

	"gorm.io/gorm"
)

func RunMigragrion(db *gorm.DB) {
	db.AutoMigrate(models.Book)
}
