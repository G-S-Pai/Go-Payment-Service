package models

import "time"

type Order struct {
	ID        string    `gorm:"type:STRING(36);primaryKey" json:"id"`
	UserID    string    `json:"user_id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}