package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Code struct {
	gorm.Model // base fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
	code       string
	expiresAt  time.Time
	status     string `gorm:"not null default:'pending' enum('pending', 'sent', 'failed')"`
}
