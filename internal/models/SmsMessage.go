package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type SmsMessage struct {
	gorm.Model
	UserId string `gorm:"not null"`
	Phone  string `gorm:"not null"`
	CodeID uint   `gorm:"not null"`
	Code   Code   `gorm:"foreignKey:CodeID"`
}

func GetSmsCodeByUserID(db *gorm.DB, userID uint) (Code, error) {
	var smsMessage SmsMessage

	if err := db.Where("user_id = ?", userID).First(&smsMessage).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return Code{}, errors.New("user not found")
		}
		return Code{}, err
	}

	var code Code
	if err := db.First(&code, smsMessage.CodeID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return Code{}, errors.New("code not found for user")
		}
		return Code{}, err
	}

	return code, nil
}

func CreateSMS(db *gorm.DB, userID string, phone string, codeID uint) error {
	smsMessage := SmsMessage{
		UserId: userID,
		Phone:  phone,
		CodeID: codeID,
	}

	if err := db.Create(&smsMessage).Error; err != nil {
		return err
	}

	return nil
}
