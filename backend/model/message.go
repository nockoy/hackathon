package model

import "time"

type Messages struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	ChannelID string    `json:"channel_id" gorm:"not null"`
	UserID    string    `json:"user_id" gorm:"not null"`
	Text      string    `json:"text" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Replies struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	ReplyToID string    `json:"reply_to_id"`
	UserID    string    `json:"user_id" gorm:"not null"`
	Text      string    `json:"text" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
