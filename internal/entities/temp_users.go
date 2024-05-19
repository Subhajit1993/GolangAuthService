package entities

import "time"

type TempUsers struct {
	ID        int       `gorm:"primaryKey;column:id" json:"id"`
	FullName  string    `gorm:"column:name" json:"full_name"`
	Active    bool      `gorm:"column:active"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (TempUsers) TableName() string {
	return "temp_users"
}
