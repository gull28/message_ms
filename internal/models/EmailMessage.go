package models

import (
	"github.com/jinzhu/gorm"
)

type EmailMessage struct {
	gorm.Model        // base fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
	UserId     string `gorm:"not null; uniqueIndex"`
	Email      string `gorm:"not null;"`
	CodeID     uint   `gorm:"not null"`
	Code       Code   `gorm:"foreignKey:CodeID"`
}

func GetEmailCodeByUserID(db *gorm.DB, userID uint) (Code, error) {
	var emailMessage EmailMessage
	if err := db.Where("user_id = ?", userID).First(&emailMessage).Error; err != nil {
		return Code{}, err
	}

	var code Code
	if err := db.Model(&emailMessage).Association("Code").Find(&code); err != nil {
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
