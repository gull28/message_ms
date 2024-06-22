package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Code struct {
	gorm.Model // base fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
	Code       string
	ExpiresAt  time.Time
	Status     string `gorm:"not null default:'pending' enum('pending', 'failed', 'verified', 'expired')"`
}

func CheckValidity(db *gorm.DB, code string) (Code, error) {
	var codeModel Code
	if err := db.Where("code = ?", code).First(&codeModel).Error; err != nil {
		return Code{}, err
	}

	return codeModel, nil
}
