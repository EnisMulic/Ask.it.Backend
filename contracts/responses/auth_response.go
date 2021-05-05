package responses

type AuthResponseModel struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

// swagger:response AuthResponse
type AuthResponse struct {
	// in: body
	Data AuthResponseModel `json:"data"`
}