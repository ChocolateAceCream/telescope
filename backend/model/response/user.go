package response

import "github.com/ChocolateAceCream/telescope/backend/model/dbmodel"

type UserInfo struct {
	Username string                 `json:"username"`
	Email    string                 `json:"email"`
	Info     map[string]interface{} `json:"info"`
}
type GoogleLoginResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	dbmodel.OAuthToken
}

type RefreshTokenResponse struct {
	RefreshToken string `json:"refresh_token"`
	ExpiredAt    int64  `json:"expired_at"`
	UserInfo
}

type GetUserResp struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Info     string `json:"info"`
}
