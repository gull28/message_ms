package models

import (
	"github.com/jinzhu/gorm"
)

type SmsMessage struct {
	gorm.Model        // base fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
	userId     any    `gorm:"not null unique index"`
	phone      string `gorm:"not null unique index "`
	code       string `gorm:"not null unique foreignkey:Code"`
}
