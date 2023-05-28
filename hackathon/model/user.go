package model

import (
	"time"
)

type Users struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"unique;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Rooms struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"unique;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Members struct {
	RoomID string `json:"id"`
	UserID string `json:"user_id"`
}

type Messages struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	ReplyToID string    `json:"reply_to_id"`
	RoomID    string    `json:"room_id" gorm:"not null"`
	From      string    `json:"from" gorm:"not null"`
	Text      string    `json:"text" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
