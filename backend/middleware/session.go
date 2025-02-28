// func AuthHandler(c *gin.Context) {
// 	user, err := utils.GetValueFromSessionByKey[db.User](c, "user")
// 	if err != nil {
// 		response.FailWithUnauthorized(c, "error.session.expired")
// 		c.Abort()
// 	}
// 	c.Set("user", user)
// }

package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ChocolateAceCream/telescope/backend/model/response"
	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/ChocolateAceCream/telescope/backend/utils"
	"github.com/gin-gonic/gin"
)

func SessionMiddleware() gin.HandlerFunc {
	return SessionHandler
}

// session expiration time should be synced between redis and cookie, but in case of network delay, just add 5 seconds buffer time to redis expiration time
func SessionHandler(c *gin.Context) {
	config := singleton.Config.Session
	// c.Cookie() return the cookie value fetched with cookieName, which is session key that used to fetch session values from redis
	// no cookie set, return unauthorized, let frontend then refresh the session
	sessionID, err := c.Cookie(config.CookieName)
	if err != nil {
		response.FailWithExpiredSession(c, "error.session.expired")
		c.Abort()
		return
	}

	ttl, err := singleton.Redis.TTL(context.TODO(), sessionID).Result()
	if err != nil {
		response.FailWithExpiredSession(c, "error.session.expired")
		c.Abort()
		return
	}
	ttlSeconds := int64(ttl.Seconds())

	// renew session in redis and cookie if session is about to expire
	if ttlSeconds < config.RefreshBeforeExpireTime {
		err := singleton.Redis.Expire(context.TODO(), sessionID, time.Duration(config.ExpireTime)*time.Second).Err()
		if err != nil {
			response.FailWithExpiredSession(c, "error.session.expired")
			c.Abort()
			return
		}

		// domain := c.Request.Host
		path := "/"
		fmt.Println("-----session middleware------ ")
		c.SetCookie(config.CookieName, sessionID, config.ExpireTime, path, "", config.Secure, config.HttpOnly)
	}

	rawSessionStr, err := singleton.Redis.Get(context.TODO(), sessionID).Result()
	if err != nil {
		response.FailWithExpiredSession(c, "error.session.expired")
		c.Abort()
		return
	}

	var session utils.Session
	json.Unmarshal([]byte(rawSessionStr), &session)

	// store session data in current context
	c.Set(config.CookieName, session)
	c.Next()

}
