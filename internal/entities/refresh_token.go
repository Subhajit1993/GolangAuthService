package entities

import "time"

type RefreshTokens struct {
	ID           int       `gorm:"primaryKey;column:id" json:"id"`
	UserId       int       `gorm:"column:user_id" json:"user_id"`
	User         Users     `gorm:"foreignKey:UserId;references:ID;constraint:OnDelete:CASCADE" json:"user"`
	RefreshToken string    `gorm:"column:refresh_token" json:"refresh_token"`
	ExpiredAt    time.Time `gorm:"column:expired_at" json:"expired_at"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

func (RefreshTokens) TableName() string {
	return "refresh_tokens"
}
