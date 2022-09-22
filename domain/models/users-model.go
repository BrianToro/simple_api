package models

import (
	"gorm.io/gorm"
	"time"
)

type Users struct {
	Id        string         `json:"id" gorm:"primaryKey"`
	UserName  string         `json:"userName"`
	PassWord  string         `json:"password"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
