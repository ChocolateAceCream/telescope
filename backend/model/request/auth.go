package request

type Claims struct {
	Email      string `json:"email"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Picture    string `json:"picture"`
	IssuedAt   int64  `json:"iat"`
	ExpiredAt  int64  `json:"exp"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
	Email        string `json:"email"`
}

type SendCodeRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required" `
	// Password string `json:"password" binding:"required,passwordCheck" ` no need for passwordCheck here
	Email string `json:"email" binding:"required,email"`
	Code  string `json:"code" binding:"required"` // email verification code
}
