package responses

// swagger:model AuthResponseModel
type AuthResponseModel struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

// swagger:model AuthResponse
type AuthResponse struct {
	// in: body
	Data AuthResponseModel `json:"data"`
}