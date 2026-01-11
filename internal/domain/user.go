package domain

import "time"

type User struct {
	ID 				uint		`gorm:"primaryKey" json:"id"`
	Email 			string		`gorm:"unique;not null" json:"email" binding:"required"`
	PasswordHash 	string		`gorm:"not null" json:"-"` // mean don't include json response
	Name 			string		`gorm:"not null" json:"name" binding:"required"`
	AvatarUrl 		*string		`json:"avatar_url,omitempty"` // mean can empty
	CreatedAt 		time.Time	`json:"created_at"`
	UpdatedAt 		time.Time	`json:"updated_at"`
}

type RegisterInput struct {
	Email 		string `json:"email" binding:"required, email"`
	Password 	string `json:"password" binding:"required, min=8"`
	Name 		string `json:"name" binding:"required, min=4"`
}

type LoginInput struct {
	Email 		string `json:"email" binding:"required, email"`
	Password 	string `json:"password" binding:"required, min=8"`
}

type LoginResponse struct {
	AccessToken string 	`json:"access_token"`
	User 		User	`json:"user"`
}