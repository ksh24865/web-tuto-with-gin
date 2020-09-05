package model

import (
	"time"
)

type Article struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Title     string    `json:"title" gorm:"type:varchar(64);"`
	Content   string    `json:"content" gorm:"type:varchar(128)"`
	CreatedAt time.Time `json:"created_at"`
}
