package database

import (
	"github.com/chowanij/go-rest-api/internal/comment"
	"github.com/jinzhu/gorm"
)

// MigrateDB - migrates our database and creates our comments table
func MigrateDB(DB *gorm.DB) error {
	if result := DB.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return result.Error
	}
	return nil
}