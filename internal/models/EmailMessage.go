package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type EmailMessage struct {
	gorm.Model
	UserId string `gorm:"not null;"`
	Email  string `gorm:"not null"`
	CodeID uint   `gorm:"not null"`
	Code   Code   `gorm:"foreignKey:CodeID"`
}

func GetEmailCodeByUserID(db *gorm.DB, userID uint) (Code, error) {
	var emailMessage EmailMessage

	if err := db.Where("user_id = ?", userID).First(&emailMessage).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return Code{}, errors.New("user not found")
		}
		return Code{}, err
	}

	var code Code
	if err := db.First(&code, emailMessage.CodeID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return Code{}, errors.New("code not found for user")
		}
		return Code{}, err
	}

	return code, nil
}

func CreateMail(db *gorm.DB, userID string, email string, codeID uint) error {
	emailMessage := EmailMessage{
		UserId: userID,
		Email:  email,
		CodeID: codeID,
	}

	if err := db.Create(&emailMessage).Error; err != nil {
		return err
	}

	return nil
}
