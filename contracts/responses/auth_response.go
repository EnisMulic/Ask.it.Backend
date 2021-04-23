package responses

// swagger:response AuthResponse
type AuthResponse struct {
	// in: body
	Data struct {
		Token string
		RefreshToken string
	}
}