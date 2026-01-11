package domain

import "time"

type RefreshToken struct {
	ID uint `gorm:"primaryKey" json:"id"`
	UserID uint `gorm:"not null;index" json:"user_id"`
	Token string `gorm:"unique;not null" json:"token"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
	
	User User `gorm:"foreignKey:UserID" json:"-"`
}

type RefreshInput struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type LogoutInput struct {
	RefreshToken 	string 	`json:"refresh_token" binding:"required"`
	UserID 			uint 	`json:"user_id"`
}