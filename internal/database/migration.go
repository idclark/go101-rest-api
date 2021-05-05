package database

import (
	"errors"
	"rest-api/internal/comment"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result.Error() != "" {
		return errors.New(result.Error())
	}
	return nil
}
