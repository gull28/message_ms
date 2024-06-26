package models

import (
	"github.com/jinzhu/gorm"
)

type SmsMessage struct {
	gorm.Model        // base fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
	UserId     string `gorm:"not null; uniqueIndex"`
	Phone      string `gorm:"not null; uniqueIndex"`
	CodeID     uint   `gorm:"not null"`
	Code       Code   `gorm:"foreignKey:CodeID"`
}

func GetSmsCodeByUserID(db *gorm.DB, userID uint) (Code, error) {
	var smsMessage SmsMessage
	if err := db.Where("user_id = ?", userID).First(&smsMessage).Error; err != nil {
		return Code{}, err
	}

	var code Code
	if err := db.Model(&smsMessage).Association("Code").Find(&code); err != nil {
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
