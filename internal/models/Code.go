package models

import (
	"time"

	"github.com/gull28/message_ms/internal/config"
	"github.com/jinzhu/gorm"
)

type Code struct {
	gorm.Model // base fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
	Code       string
	UserId     uint
	ExpiresAt  time.Time
	Attempt    int
	Status     string `gorm:"not null default:'pending' enum('pending', 'failed', 'verified', 'expired')"`
}

func CheckValidity(db *gorm.DB, code string) (Code, error) {
	var codeModel Code
	if err := db.Where("code = ?", code).First(&codeModel).Error; err != nil {
		return Code{}, err
	}

	return codeModel, nil
}

func GetAttemptCount(db *gorm.DB, userId string) (int, error) {
	var code Code
	var count int

	resendTimer := config.LoadConfig().CodeSettings.ResendTimer

	timeThreshold := time.Now().Add(-time.Duration(resendTimer) * time.Minute)

	if err := db.Model(&code).Where("UserId = ? AND CreatedAt > ?", userId, timeThreshold).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
