package models

import (
	"time"
)

type User struct {
    ID uint `gorm:"primaryKey" json:"id"` 
    FirstName string `gorm:"type:varchar(255);not null" json:"firstName"`
    LastName string `gorm:"type:varchar(255);not null" json:"lastName"`
    Email string `gorm:"uniqueIndex;notnull" json:"email"`
    Password string `gorm:"not null" json:"password"`
    CreatedAt time.Time `gorm:"not null" json:"createdAt"`
    UpdatedAt time.Time `gorm:"not null" json:"UpdatedAt"`
}
