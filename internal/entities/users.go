package entities

import "time"

type Users struct {
	ID                 int       `gorm:"primaryKey;column:id" json:"id"`
	FullName           string    `gorm:"column:name" json:"full_name"`
	Email              string    `gorm:"column:email"`
	Active             bool      `gorm:"column:active"`
	CreatedAt          time.Time `gorm:"column:created_at"`
	UpdatedAt          time.Time `gorm:"column:updated_at"`
	Verified           bool      `gorm:"column:verified"`
	RegistrationSource string    `gorm:"column:registration_source"`
}

func (Users) TableName() string {
	return "users"
}
