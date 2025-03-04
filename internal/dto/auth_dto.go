package dto

type LoginRequestDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponseDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshRequestDTO struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponseDTO struct {
	AccessToken string `json:"access_token"`
}
