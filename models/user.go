package models

import "time"

type User struct {
	Id         uint      `gorm:"primaryKey;autoIncrement"`
	Fullname   string    `gorm:"size:100;not null"`
	Email      string    `gorm:"size:150;not null"`
	Password   string    `gorm:"not null"`
	Active     bool      `gorm:"default:false"`
	CreatedBy  string    `gorm:"size:100;not null"`
	CreatedAt  time.Time `gorm:"not null"`
	ModifiedBy *string   `gorm:"size:100"`
	ModifiedAt *time.Time
}
