package service

import (
	"encoding/json"
	"fmt"
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

	// check whitelist
	whitelisted, err := utils.IsEmailWhitelisted(claims.Email)
	if err != nil {
		return
	}

	if !whitelisted {
		err = fmt.Errorf("error.email.not.whitelisted")
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

	err = utils.RenewSession(c, user)
	if err != nil {
		return
	}

	// utils.DeleteSession(c, utils.RefreshTokenCookieName)
	utils.SetRefreshToken(c, user.Email)
	fmt.Println("------final---")
	cookie, err := c.Cookie(utils.RefreshTokenCookieName)
	fmt.Println("RefreshTokenCookieName: ", cookie)
	cookie, err = c.Cookie(singleton.Config.Session.CookieName)
	fmt.Println("Session cookie: ", cookie)

	return
}

func (AuthService *AuthService) SendCode(c *gin.Context, email string) (err error) {
	randomCode := utils.RandomNumber(singleton.Config.Captcha.Length)
	singleton.Redis.Set(c, singleton.Config.Email.Prefix+":"+email, randomCode, time.Duration(singleton.Config.Email.Expiration)*time.Minute)
	body := fmt.Sprintf("verification code is <b>%v</b>, expired in %v minutes", randomCode, singleton.Config.Email.Expiration)
	return utils.SendMail(email, "Verification Code", body)
}

func (AuthService *AuthService) Register(c *gin.Context, payload request.RegisterRequest) (user db.User, err error) {
	// check email whitelist
	whitelisted, err := utils.IsEmailWhitelisted(payload.Email)
	if err != nil {
		singleton.Logger.Error("check email whitelist failed", zap.Error(err))
		return
	}
	if !whitelisted {
		err = fmt.Errorf("error.email.not.whitelisted")
		return
	}
	// check if email exists
	_, err = userDao.GetUserByEmail(c, payload.Email)
	if err == nil {
		err = fmt.Errorf("error.email.exists")
		return
	}

	// check if code and email matched
	code, err := singleton.Redis.Get(c, singleton.Config.Email.Prefix+":"+payload.Email).Result()
	if err != nil {
		err = fmt.Errorf("error.invalid.code")
		return
	}
	if code != payload.Code {
		err = fmt.Errorf("error.invalid.code")
		return
	}

	// start a transaction to create user and password login
	tx, err := singleton.DB.BeginTx(c, pgx.TxOptions{})
	if err != nil {
		singleton.Logger.Error("begin tx failed", zap.Error(err))
		err = fmt.Errorf("error.failed.operation")
		return
	}
	defer tx.Rollback(c)
	utils.WithTx(c, tx)

	createUserPayload := db.CreateNewUserParams{
		Email:    payload.Email,
		Username: payload.Username,
		Info:     []byte("{}"),
	}
	err = userDao.CreateUser(c, createUserPayload)
	if err != nil {
		err = fmt.Errorf("error.failed.operation")
		return
	}

	passwordLoginPayload := db.CreateNewPasswordLoginParams{
		Email:    payload.Email,
		Password: payload.Password,
	}
	err = userDao.CreateNewPasswordLogin(c, passwordLoginPayload)
	if err != nil {
		err = fmt.Errorf("error.failed.operation")
		return
	}

	err = tx.Commit(c)
	if err != nil {
		singleton.Logger.Error("commit tx failed", zap.Error(err))
		err = fmt.Errorf("error.failed.operation")
		return
	}

	// create session
	user, err = userDao.GetUserByEmail(c, payload.Email)
	if err != nil {
		err = fmt.Errorf("error.failed.operation")
		return
	}

	// login success, set session
	err = utils.NewSession(c, user)
	if err != nil {
		err = fmt.Errorf("error.failed.operation")
		return
	}

	err = utils.SetRefreshToken(c, user.Email)
	if err != nil {
		err = fmt.Errorf("error.failed.operation")
		return
	}
	return
}
