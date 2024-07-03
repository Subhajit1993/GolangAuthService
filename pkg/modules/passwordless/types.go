package passwordless

import (
	"time"
)

type PublicProfile struct {
	ID                 int
	Email              string `json:"email"`
	DisplayName        string `json:"nickname"`
	RegistrationSource string `json:"sub"`
	Picture            string `json:"picture"`
	FullName           string `json:"name"` // This is nickname from openid
	Verified           bool   `json:"email_verified"`
}

type passwordlessRegistration struct {
	UserId      int
	Id          int       `json:"id"`
	UserHandle  string    `json:"user_handle"`
	RPID        string    `json:"rpid"`
	Challenge   string    `json:"credential" gorm:"column:credential"`
	DisplayName string    `json:"display_name"`
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	ExpiredAt   time.Time `json:"expired_at"`
}
