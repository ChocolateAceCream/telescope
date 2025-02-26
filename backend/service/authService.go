package service

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	db "github.com/ChocolateAceCream/telescope/backend/db/sqlc"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/ChocolateAceCream/telescope/backend/model/dbmodel"
	"github.com/ChocolateAceCream/telescope/backend/model/request"
	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/ChocolateAceCream/telescope/backend/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthService struct{}

func (a *AuthService) Login(c *gin.Context, payload request.LoginRequest) (user db.User, err error) {
	err = userDao.VerifyUserCredentials(c, payload.Email, payload.Password)
	if err != nil {
		return
	}
	user, err = userDao.GetUserByEmail(c, payload.Email)
	if err != nil {
		return
	}

	err = utils.NewSession(c, user)
	if err != nil {
		return
	}

	err = utils.SetRefreshToken(c, user.Email)
	return
}

func (authService *AuthService) ExchangeCodeForToken(c *gin.Context, code string) (err error) {
	data := url.Values{}
	data.Set("client_id", singleton.Config.OAuth.Google.ClientID)
	data.Set("client_secret", singleton.Config.OAuth.Google.ClientSecret)
	data.Set("code", code)
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", singleton.Config.OAuth.Google.RedirectURL)

	resp, err := http.PostForm("https://oauth2.googleapis.com/token", data)
	if err != nil {
		singleton.Logger.Error("Failed to exchange code for token", zap.Error(err))
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		singleton.Logger.Error("Failed to read exchange code request body", zap.Error(err))
		return
	}

	token := dbmodel.OAuthToken{}

	err = json.Unmarshal(body, &token)
	if err != nil {
		singleton.Logger.Error("Failed to unmarshal exchange code request body", zap.Error(err))
		return
	}
	claimsJSON, err := utils.Decoder(token.IdToken)
	if err != nil {
		singleton.Logger.Error("Failed to decode id token", zap.Error(err))
		return
	}
	var claims request.Claims
	err = json.Unmarshal(claimsJSON, &claims)
	if err != nil {
		singleton.Logger.Error("Failed to unmarshal claims", zap.Error(err))
		return
	}

	tx, err := singleton.DB.BeginTx(c, pgx.TxOptions{})
	if err != nil {
		singleton.Logger.Error("begin tx failed", zap.Error(err))
		return
	}
	defer tx.Rollback(c)
	utils.WithTx(c, tx)

	userPayload := map[string]interface{}{
		"given_name":  claims.GivenName,
		"family_name": claims.FamilyName,
		"picture":     claims.Picture,
	}
	info, _ := json.Marshal(userPayload)

	createUserPayload := db.CreateNewUserParams{
		Email:    claims.Email,
		Username: claims.Name,
		Info:     info,
	}
	err = userDao.CreateUser(c, createUserPayload)
	if err != nil {
		return
	}

	googleLoginPayload := db.GoogleLoginParams{
		Email:       claims.Email,
		Username:    claims.Name,
		AccessToken: token.AccessToken,
		IssuedAt:    pgtype.Timestamp{Time: time.Unix(claims.IssuedAt, 0), Valid: true},
		ExpiredAt:   pgtype.Timestamp{Time: time.Unix(claims.ExpiredAt, 0), Valid: true},
	}

	err = userDao.CreateGoogleLogin(c, googleLoginPayload)
	if err != nil {
		return
	}

	err = tx.Commit(c)
	if err != nil {
		singleton.Logger.Error("commit tx failed", zap.Error(err))
		return
	}

	user, err := userDao.GetUserByEmail(c, claims.Email)
	if err != nil {
		return
	}

	// login success, set session
	err = utils.NewSession(c, user)
	if err != nil {
		return
	}

	err = utils.SetRefreshToken(c, user.Email)
	if err != nil {
		return
	}
	return
	// singleton.Logger.Info("Successfully exchanged code for token", zap.Any("token", token))
}

// 1. get user info; 2. create new session; 3. create new refresh token
func (AuthService *AuthService) RefreshToken(c *gin.Context) (err error) {
	refreshToken, err := c.Cookie(utils.RefreshTokenCookieName)
	if err != nil {
		return
	}
	email, err := singleton.Redis.Get(c, refreshToken).Result()
	if err != nil {
		return
	}
	user, err := userDao.GetUserByEmail(c, email)
	if err != nil {
		return
	}

	err = utils.NewSession(c, user)
	if err != nil {
		return
	}

	err = utils.SetRefreshToken(c, user.Email)
	return
}
