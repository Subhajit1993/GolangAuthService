package entities

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type PasswordlessAuthStatus string

const (
	REGISTRATION_IN_PROGRESS PasswordlessAuthStatus = "REGISTRATION_IN_PROGRESS"
	REGISTRATION_SUCCESS     PasswordlessAuthStatus = "REGISTRATION_SUCCESS"
	REGISTRATION_FAILED      PasswordlessAuthStatus = "REGISTRATION_FAILED"
	LOGIN_SUCCESS            PasswordlessAuthStatus = "LOGIN_SUCCESS"
)

type Passwordless struct {
	ID                             int                    `gorm:"primaryKey;column:id" json:"id"`
	UserId                         int                    `gorm:"column:user_id" json:"user_id"`
	Active                         bool                   `gorm:"column:active" json:"active"`
	IsDeleted                      bool                   `gorm:"column:is_deleted" json:"is_deleted"`
	User                           Users                  `gorm:"foreignKey:UserId;references:ID" json:"user"`
	UserHandle                     string                 `gorm:"column:user_handle" json:"user_handle" validate:"required"` //A user-specific identifier that links the WebAuthn credential to the user account
	RPID                           string                 `gorm:"column:rpid" json:"rpid"`                                   //The relying party identifier
	Credential                     string                 `gorm:"column:credential" json:"challenge" validate:"required"`
	CredentialType                 string                 `gorm:"column:credential_type" json:"credential_type" validate:"required"` //The type of the credential, such as "platform" for built-in authenticators or "cross-platform" for external authenticators like security keys
	Counter                        int                    `gorm:"column:counter;not null" json:"counter"`
	Status                         PasswordlessAuthStatus `gorm:"column:status" json:"status"`
	ExpiredAt                      time.Time              `gorm:"column:expired_at;type:timestamp" json:"expired_at"`
	RawRegistrationAttestationData string                 `gorm:"column:raw_registration_attestation_data" json:"raw_registration_attestation_data"`
	RawSessionData                 string                 `gorm:"column:raw_session_data" json:"raw_session_data"`
	PublicKey                      string                 `gorm:"column:public_key" json:"public_key"` // This we will get during finish registration
	CreatedAt                      time.Time              `gorm:"column:created_at"`
	UpdatedAt                      time.Time              `gorm:"column:updated_at"`
}

func (Passwordless) TableName() string {
	return "passwordless"
}

func (p Passwordless) Validate() error {
	validate := validator.New()
	err := validate.Struct(p)
	if err != nil {
		return err
	}
	return nil
}
