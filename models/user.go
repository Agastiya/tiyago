package models

import (
	"time"
)

type User struct {
	Id         int64     `gorm:"primaryKey;autoIncrement"`
	Fullname   string    `gorm:"size:100;not null"`
	Username   string    `gorm:"size:15;uniqueIndex"`
	Email      string    `gorm:"size:150;uniqueIndex;not null"`
	Password   string    `gorm:"not null"`
	Active     bool      `gorm:"default:false"`
	CreatedBy  string    `gorm:"size:100;not null"`
	CreatedAt  time.Time `gorm:"not null"`
	ModifiedBy *string   `gorm:"size:100"`
	DeletedBy  *string   `gorm:"size:100"`
	ModifiedAt *time.Time
	DeletedAt  *time.Time
}
