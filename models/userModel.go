package models

import (
	"time"
)

type User struct {
    ID uint `gorm:"primaryKey" json:"id"` 
    FirstName string `gorm:"type:varchar(255);not null" json:"firstName"`
    LastName string `gorm:"type:varchar(255);not null" json:"lastName"`
    Email string `gorm:"uniqueIndex;not null" json:"email"`
    Password string `gorm:"not null" json:"password"`
    Phone string `gorm:"type:varchar(255);not null" json:"phone"`
    CreatedAt time.Time `gorm:"not null" json:"createdAt"`
    UpdatedAt time.Time `gorm:"not null" json:"updatedAt"`
    Students []Student `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Student struct {
    ID uint `gorm:"primaryKey" json:"id"` 
    FirstName string `gorm:"type:varchar(255);not null" json:"firstName"`
    LastName string `gorm:"type:varchar(255);not null" json:"lastName"`
    Email string `json:"email"`
    Password string `gorm:"not null" json:"password"`
    Phone string `gorm:"type:varchar(255);not null" json:"phone"`
    Address string `gorm:"type:varchar(255);not null" json:"address"`
    UserID uint `json:"userID"`
    CreatedAt time.Time `gorm:"not null" json:"createdAt"`
    UpdatedAt time.Time `gorm:"not null" json:"updatedAt"`
}
