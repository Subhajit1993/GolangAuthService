package openid

type PasswordlessRegistrationBeginAPIRequest struct {
	Email       string `json:"email" binding:"required"`
	DisplayName string `json:"display_name" binding:"required"`
}
