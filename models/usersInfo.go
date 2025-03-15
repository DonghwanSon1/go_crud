package models

import (
	"gorm.io/gorm"
	"time"
)

type UsersInfo struct {
	UserId    string         `gorm:"primaryKey;type:varchar(255);not null" json:"user_id"`
	Password  string         `gorm:"type:varchar(255);not null" json:"password"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	Phone     string         `gorm:"type:varchar(32)" json:"phone"`
	Email     string         `gorm:"type:varchar(255)" json:"email"`
	About     string         `gorm:"type:varchar(1024)" json:"about"`
	Gender    string         `gorm:"type:varchar(255)" json:"gender"`
	Birth     time.Time      `gorm:"type:date" json:"birth"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
