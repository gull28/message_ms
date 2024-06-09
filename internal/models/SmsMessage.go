package models

import (
	"github.com/jinzhu/gorm"
)

type SmsMessage struct {
	gorm.Model        // base fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
	UserId     uint   `gorm:"not null; uniqueIndex"`
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
