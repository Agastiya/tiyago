package dto

type LoginByEmailRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Id           int64  `json:"id"`
	Fullname     string `json:"fullname"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
