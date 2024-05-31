package models

import (
	"github.com/jinzhu/gorm"
)

type EmailMessage struct {
	gorm.Model        // base fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
	userId     any    `gorm:"not null unique index"`
	email      string `gorm:"not null unique index"`
	code       string `gorm:not null unique foreignkey:Code` // foreign key to Code
}
