package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	db "github.com/ChocolateAceCream/telescope/backend/db/sqlc"
	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const RefreshTokenCookieName = "refresh-token"

type Session struct {
	Content map[string]interface{} `json:"content"`
	UUID    string                 `json:"uuid"`
	Lock    *sync.Mutex            `json:"lock"`
}

func SetRefreshToken(c *gin.Context, email string) (err error) {
	config := singleton.Config.Session
	UUID := uuid.New().String()
	// domain := c.Request.Host
	path := "/"
	fmt.Println("--set refresh----request host:------ ")
	c.SetCookie(RefreshTokenCookieName, UUID, config.RefreshTokenExpireTime, path, "", config.Secure, config.HttpOnly)

	err = singleton.Redis.Set(context.TODO(), UUID, email, time.Duration(config.RefreshTokenExpireTime)*time.Second).Err()
	if err != nil {
		singleton.Logger.Error("set refresh token failed", zap.Error(err))
	}
	return
}

// renew session in cookie and redis
func RenewSession(c *gin.Context, user db.User) (err error) {
	config := singleton.Config.Session
	sessionID, err := c.Cookie(config.CookieName)
	if err != nil {
		// if cookie not found, create new session
		err = NewSession(c, user)
		return
	}
	_, err = singleton.Redis.Get(context.TODO(), sessionID).Result()
	if err != nil {
		// cannot find session in redis, create new session
		singleton.Logger.Error("get session failed", zap.Error(err))

		path := "/"
		c.SetCookie(config.CookieName, sessionID, config.ExpireTime, path, "", config.Secure, config.HttpOnly)

		session := Session{
			Content: map[string]interface{}{
				"user": user,
			},
			UUID: sessionID,
			Lock: &sync.Mutex{},
		}
		jsonStr, _ := json.Marshal(session)

		err = singleton.Redis.Set(context.TODO(), sessionID, jsonStr, time.Duration(config.ExpireTime)*time.Second).Err()
		if err != nil {
			singleton.Logger.Error("set session failed", zap.Error(err))
			return err
		}
	}
	// renew session expire time
	err = singleton.Redis.Expire(context.TODO(), sessionID, time.Duration(config.ExpireTime)*time.Second).Err()
	return
}

// set new session in cookie and redis
func NewSession(c *gin.Context, user db.User) (err error) {
	config := singleton.Config.Session
	UUID := uuid.New().String()
	// domain := c.Request.Host
	path := "/"
	fmt.Println("--newsession------ ")
	c.SetCookie(config.CookieName, UUID, config.ExpireTime, path, "", config.Secure, config.HttpOnly)

	session := Session{
		Content: map[string]interface{}{
			"user": user,
		},
		UUID: UUID,
		Lock: &sync.Mutex{},
	}
	jsonStr, err := json.Marshal(session)
	if err != nil {
		singleton.Logger.Error("marshal session failed", zap.Error(err))
		return
	}
	err = singleton.Redis.Set(context.TODO(), UUID, jsonStr, time.Duration(config.ExpireTime)*time.Second).Err()
	if err != nil {
		singleton.Logger.Error("set session failed", zap.Error(err))
	}
	return
}

// obtain val from session by given key
func (s *Session) Get(key string) (val interface{}, err error) {
	raw, err := singleton.Redis.Get(context.TODO(), s.UUID).Result()
	if err != nil {
		singleton.Logger.Error("get session failed", zap.Error(err))
		return
	}
	var session Session
	err = json.Unmarshal([]byte(raw), &session)
	if err != nil {
		singleton.Logger.Error("unmarshal session failed", zap.Error(err))
		return
	}
	if val, ok := session.Content[key]; ok {
		return val, nil
	}
	err = errors.New("not found key: " + key)
	return
}

// renew session expire time
/*
func (s *Session) Renew(c *gin.Context) {
	newExpire := time.Now().Unix() + int64(singleton.Config.Session.ExpireTime)
	jsonStr, _ := json.Marshal(s)
	singleton.Redis.Set(c, s.UUID, jsonStr, time.Duration(newExpire))
}
*/

// set value in session by given key, also renew the session expire time
func (s *Session) SetValueByKey(key string, val any) (err error) {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	raw, err := singleton.Redis.Get(context.TODO(), s.UUID).Result()
	if err != nil {
		singleton.Logger.Error("get session failed", zap.Error(err))
		return
	}
	var session Session
	err = json.Unmarshal([]byte(raw), &session)
	if err != nil {
		singleton.Logger.Error("unmarshal session failed", zap.Error(err))
		return
	}
	session.Content[key] = val
	updated, err := json.Marshal(session)
	if err != nil {
		singleton.Logger.Error("marshal updated session failed", zap.Error(err))
		return
	}
	newExpire := time.Duration(time.Now().Unix()+int64(singleton.Config.Session.ExpireTime)) * time.Second
	singleton.Redis.Set(context.TODO(), s.UUID, updated, newExpire)
	return
}

func GetSession(c *gin.Context) *Session {
	cookie, ok := c.Get(singleton.Config.Session.CookieName)
	if !ok {
		singleton.Logger.Error("cannot retrieve cookie from current context")
		return nil
	}
	session, ok := cookie.(Session)
	if !ok {
		singleton.Logger.Error("cookie is not of type Session")
		return nil
	}
	return &session
}

func GetValueFromSessionByKey[T any](c *gin.Context, key string) (val T, err error) {
	session, ok := c.Get(singleton.Config.Session.CookieName)
	if !ok {
		singleton.Logger.Error("cannot retrieve session from current context")
	}
	content := session.(Session).Content
	raw, ok := content[key]
	if !ok {
		singleton.Logger.Error("cannot retrieve value from session by given key")
		return
	}

	jsonStr, err := json.Marshal(raw)
	if err != nil {
		singleton.Logger.Error("marshal failed", zap.Error(err))
		return
	}
	err = json.Unmarshal(jsonStr, &val)

	if err != nil {
		singleton.Logger.Error("cannot retrieve value from session by given key")
	}
	return
}

func DeleteSession(c *gin.Context, cookieName string) (err error) {
	config := singleton.Config.Session
	cookie, err := c.Cookie(cookieName)
	fmt.Println("---------cookie:---------- ", cookieName)
	if err != nil {
		singleton.Logger.Error("cannot retrieve cookie from current context", zap.Error(err))
		return
	}
	err = singleton.Redis.Del(context.TODO(), cookie).Err()
	if err != nil {
		singleton.Logger.Error("delete session failed", zap.Error(err))
		return
	}
	fmt.Println("--del----request host:------ ", c.Request.Host)
	c.SetCookie(config.CookieName, "", -1, "/", "", config.Secure, config.HttpOnly)
	return
}

/*
func GetSessionRemainingDuration(s *Session) (duration time.Duration, err error) {
	remain := s.ExpireTime - time.Now().Unix()
	if remain < 0 {
		singleton.Logger.Error("session has expired")
		return duration, errors.New("session has expired")
	}
	duration = time.Duration(remain) * time.Second
	return duration, nil
}
*/
