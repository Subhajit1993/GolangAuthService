package entities

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type Passwordless struct {
	ID              int       `gorm:"primaryKey;column:id" json:"id"`
	UserId          int       `gorm:"column:user_id" json:"user_id"`
	User            Users     `gorm:"foreignKey:UserId;references:ID" json:"user"`
	UserHandle      string    `gorm:"column:user_handle" json:"user_handle" validate:"required"` //A user-specific identifier that links the WebAuthn credential to the user account
	RPID            string    `gorm:"column:rpid" json:"rpid"`                                   //The relying party identifier
	Challenge       string    `gorm:"column:credential" json:"challenge" validate:"required"`
	CredentialType  string    `gorm:"column:credential_type" json:"credential_type" validate:"required"` //The type of the credential, such as "platform" for built-in authenticators or "cross-platform" for external authenticators like security keys
	Counter         int       `gorm:"column:counter;not null" json:"counter"`
	AttestationData string    `gorm:"column:attestation_data" json:"attestation_data"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
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
