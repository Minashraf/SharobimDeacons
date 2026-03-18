package payload

type RefreshPayload struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
