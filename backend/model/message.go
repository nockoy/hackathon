package model

import "time"

type Messages struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	ReplyToID string    `json:"reply_to_id"`
	RoomID    string    `json:"room_id" gorm:"not null"`
	UserID    string    `json:"user_id" gorm:"not null"`
	Text      string    `json:"text" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
