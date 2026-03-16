package payload

type RefreshPayload struct {
	UserID       int64  `json:"user_id" binding:"required"`
	RefreshToken string `json:"refresh_token" binding:"required"`
}
