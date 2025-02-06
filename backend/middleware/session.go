package middleware

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/ChocolateAceCream/telescope/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SessionMiddleware() gin.HandlerFunc {
	return SessionHandler
}

func SessionHandler(c *gin.Context) {
	config := singleton.Config.Session
	// c.Cookie() return the cookie value fetched with cookieName, which is session key that used to fetch session values from redis
	// if cookie has been set, continue, otherwise create new session
	if sessionID, err := c.Cookie(config.CookieName); err == nil {
		// if session not expired in redis, store session in current context, otherwise create new session
		if rawSessionStr, err := singleton.Redis.Get(context.TODO(), sessionID).Result(); err == nil {
			var session utils.Session
			json.Unmarshal([]byte(rawSessionStr), &session)
			if (session.ExpireTime - time.Now().Unix()) < config.RefreshBeforeExpireTime {
				session.Renew(c)
				c.SetCookie(config.CookieName, session.UUID, int(session.ExpireTime-time.Now().Unix()), "/", c.Request.Host, config.Secure, config.HttpOnly)
			}
			c.Set(config.Key, session)
			return
		}
	}
	UUID := uuid.New().String()
	domain := c.Request.Host
	path := "/"
	c.SetCookie(config.CookieName, UUID, config.ExpireTime, path, domain, config.Secure, config.HttpOnly)
	newSession := utils.Session{
		Cookie:     config.CookieName,
		ExpireTime: time.Now().Unix() + int64(config.ExpireTime),
		Content:    make(map[string]interface{}),
		UUID:       UUID,
		Lock:       &sync.Mutex{},
	}
	c.Set(config.Key, newSession)
	jsonStr, _ := json.Marshal(newSession)
	singleton.Redis.Set(c, UUID, jsonStr, time.Duration(config.ExpireTime)*time.Second)
}
