package models

import (
	"errors"
	"time"

	"github.com/gull28/message_ms/internal/config"
	"github.com/jinzhu/gorm"
)

type Code struct {
	gorm.Model // base fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
	Code       string
	UserId     string
	ExpiresAt  time.Time
	Attempt    int
	Status     string `gorm:"not null default:'pending' enum('pending', 'failed', 'verified')"`
}

func CheckValidity(db *gorm.DB, code string, userId string) (bool, error) {
	allowedAttempts := config.LoadConfig().CodeSettings.Attempts

	var codeModel Code
	if err := db.Where("UserId = ?", userId).Order("CreatedAt DESC").First(&codeModel).Error; err != nil {
		return false, err
	}

	if codeModel.Attempt >= allowedAttempts {
		err := errors.New("Maximum attempts reached")
		return false, err
	}

	if codeModel.Code != code {
		err := errors.New("Invalid code")

		codeModel.Attempt++

		if err := db.Save(&codeModel).Error; err != nil {
			return false, err
		}

		return false, err
	}

	return true, nil
}

// GetResendCount returns the number of resend attempts for a given user within a specified time frame
func GetResendCount(db *gorm.DB, userId string) (int, error) {
	var code Code
	var count int

	resendTimer := config.LoadConfig().CodeSettings.ResendTimer

	timeThreshold := time.Now().Add(-time.Duration(resendTimer) * time.Minute)

	if err := db.Model(&code).Where("UserId = ? AND CreatedAt > ?", userId, timeThreshold).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func CreateCode(db *gorm.DB, userId string, code string, expiresAt time.Time) (Code, error) {
	codeModel := Code{
		Code:      code,
		UserId:    userId,
		Attempt:   0,
		ExpiresAt: expiresAt,
		Status:    "pending",
	}

	if err := db.Create(&codeModel).Error; err != nil {
		return Code{}, err
	}

	return codeModel, nil
}
