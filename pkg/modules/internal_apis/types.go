package internal_apis

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	ID       int
}
