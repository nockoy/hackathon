package model

import "time"

type Messages struct {
	MessageID string    `json:"id" gorm:"primaryKey"`
	ChannelID string    `json:"channel_id" gorm:"not null"`
	UserID    string    `json:"user_id" gorm:"not null"`
	Text      string    `json:"text" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MessagesAndUserInfo struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	ChannelID string    `json:"channel_id" gorm:"not null"`
	UserID    string    `json:"user_id" gorm:"not null"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Icon      string    `json:"icon"`
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

type RepliesAndUserInfo struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	ReplyToID string    `json:"reply_to_id" gorm:"not null"`
	UserID    string    `json:"user_id" gorm:"not null"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Icon      string    `json:"icon"`
	Text      string    `json:"text" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
