package openid

type Profile struct {
	Email              string `json:"email"`
	DisplayName        string `json:"nickname"`
	RegistrationSource string `json:"sub"`
	Picture            string `json:"picture"`
	FullName           string `json:"name"` // This is nickname from openid
	Verified           bool   `json:"email_verified"`
}
