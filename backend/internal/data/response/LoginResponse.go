package response

type LoginResponse struct {
	Token        string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	RefreshToken string `json:"refresh_token" example:"GciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	UserId       int64  `json:"user_id" example:"1"`
}
