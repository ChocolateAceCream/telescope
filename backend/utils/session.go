package utils

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Session struct {
	Cookie     string                 `json:"cookie"`
	ExpireTime int64                  `json:"expireTime"`
	Content    map[string]interface{} `json:"content"`
	UUID       string                 `json:"uuid"`
	Lock       *sync.Mutex            `json:"lock"`
}

// obtain val from session by given key
func (s *Session) Get(key string) (val interface{}, err error) {
	raw, err := singleton.Redis.Get(context.TODO(), s.UUID).Result()
	if err != nil {
		return
	}
	var session Session
	err = json.Unmarshal([]byte(raw), &session)
	if err != nil {
		return
	}
	if val, ok := session.Content[key]; ok {
		return val, nil
	}
	err = errors.New("not found key: " + key)
	return
}

// renew session expire time
func (s *Session) Renew(c *gin.Context) {
	newExpire := time.Now().Unix() + int64(singleton.Config.Session.ExpireTime)
	jsonStr, _ := json.Marshal(s)
	singleton.Redis.Set(c, s.UUID, jsonStr, time.Duration(newExpire))
}

func (s *Session) Set(key string, val any) (err error) {
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
	newExpire, err := GetSessionRemainingDuration(s)
	if err != nil {
		return
	}
	singleton.Redis.Set(context.TODO(), s.UUID, updated, newExpire)
	return
}

func GetSession(c *gin.Context) *Session {
	cookie, ok := c.Get(singleton.Config.Session.Key)
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
	session := GetSession(c)
	if session == nil {
		singleton.Logger.Error("no session found")
		return val, errors.New("no session found")
	}
	raw, err := session.Get(key)
	if err != nil {
		return
	}
	jsonStr, err := json.Marshal(raw)
	if err != nil {
		singleton.Logger.Error("marshal failed", zap.Error(err))
		return
	}
	err = json.Unmarshal(jsonStr, &val)
	return
}

func GetSessionRemainingDuration(s *Session) (duration time.Duration, err error) {
	remain := s.ExpireTime - time.Now().Unix()
	if remain < 0 {
		singleton.Logger.Error("session has expired")
		return duration, errors.New("session has expired")
	}
	duration = time.Duration(remain) * time.Second
	return duration, nil
}
