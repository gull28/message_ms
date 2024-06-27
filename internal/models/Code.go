package models

import (
	"errors"
	"time"

	"github.com/gull28/message_ms/internal/config"
	"github.com/jinzhu/gorm"
)

// Code represents the structure for code management
type Code struct {
	gorm.Model // base fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
	Code       string
	UserId     string
	ExpiresAt  time.Time
	Attempt    int
	Status     string `gorm:"not null;default:'pending';type:ENUM('pending', 'failed', 'verified')"`
}

func CheckValidity(db *gorm.DB, code string, userId string) (bool, error) {
	allowedAttempts := config.LoadConfig().CodeSettings.Attempts

	var codeModel Code

	err := db.Where("user_id = ? AND status = ?", userId, "pending").
		Order("created_at DESC").
		First(&codeModel).Error
	if err != nil {
		return false, err
	}

	if codeModel.Status == "verified" {
		return false, errors.New("code already verified")
	}

	if codeModel.Attempt >= allowedAttempts {
		if codeModel.Status != "failed" {
			codeModel.Status = "failed"
			if err := db.Save(&codeModel).Error; err != nil {
				return false, err
			}
		}
		return false, errors.New("maximum attempts reached")
	}

	if codeModel.Code != code {
		codeModel.Attempt++
		if err := db.Save(&codeModel).Error; err != nil {
			return false, err
		}
		return false, errors.New("invalid code")
	}

	codeModel.Status = "verified"
	if err := db.Save(&codeModel).Error; err != nil {
		return false, err
	}

	return true, nil
}

func GetResendCount(db *gorm.DB, userId string) (int, error) {
	var count int
	resendTimer := config.LoadConfig().CodeSettings.ResendTimer
	timeThreshold := time.Now().Add(-time.Duration(resendTimer) * time.Minute)

	err := db.Model(&Code{}).Where("user_id = ? AND created_at > ?", userId, timeThreshold).Count(&count).Error
	if err != nil {
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
