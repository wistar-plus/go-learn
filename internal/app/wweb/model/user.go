package model

import (
	"time"
)

type User struct {
	Uid       uint32    `gorm:"primary_key;auto_increment" json:"uid"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Pwd       string    `gorm:"size:100;not null;" json:"pwd"`
	NickName  string    `gorm:"size:255;not null" json:"nickname"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type UserLoginParam struct {
	Email string `json:"email"`
	Pwd   string `json:"pwd"`
}
