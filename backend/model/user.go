package model

import (
	"time"
)

type Users struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Members struct {
	WorkspaceID string `json:"workspace_id"`
	UserID      string `json:"user_id"`
}
