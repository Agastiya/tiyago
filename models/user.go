package models

import "time"

type User struct {
	ID          uint `gorm:"primaryKey"`
	Fullname    string
	Email       string
	Password    string
	Status      string
	CreatedBy   string
	CreatedTime *time.Time
}
