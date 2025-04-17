package utils

import (
	"context"

	"github.com/ChocolateAceCream/telescope/backend/singleton"
	"go.uber.org/zap"
)

const RedisEmailWhitelistKey = "email:whitelist"

func InitEmailWhitelist() (err error) {
	c := context.Background()
	// check if the email whitelist has been initialized
	exists, err := singleton.Redis.Exists(c, RedisEmailWhitelistKey).Result()
	if err != nil {
		singleton.Logger.Error("check email whitelist failed", zap.Error(err))
		return
	}
	if exists == 1 {
		// already initialized
		singleton.Logger.Info("email whitelist already initialized")
		return
	}
	// initialize the email whitelist
	emails, err := singleton.Query.GetEmailWhitelist(c)
	if err != nil {
		return
	}
	_, err = singleton.Redis.SAdd(c, RedisEmailWhitelistKey, SliceToInterfaceSlice(emails)...).Result()
	return
}

func IsEmailWhitelisted(email string) (result bool, err error) {
	c := context.Background()
	result, err = singleton.Redis.SIsMember(c, RedisEmailWhitelistKey, email).Result()
	if err != nil {
		singleton.Logger.Error("check email whitelist failed", zap.Error(err))
	}
	return
}

// first add to db, then to redis
func AddEmailToWhitelist(email string) (err error) {
	c := context.Background()
	err = singleton.Query.AddEmailWhitelist(c, email)
	if err != nil {
		singleton.Logger.Error("add email to whitelist table failed", zap.Error(err))
		return
	}
	_, err = singleton.Redis.SAdd(c, RedisEmailWhitelistKey, email).Result()
	if err != nil {
		singleton.Logger.Error("add email to redis whitelist set failed", zap.Error(err))
	}
	return
}

// first remove from redis, then remove from db
func RemoveEmailFromWhitelist(email string) (err error) {
	c := context.Background()
	_, err = singleton.Redis.SRem(c, RedisEmailWhitelistKey, email).Result()
	if err != nil {
		singleton.Logger.Error("remove email from whitelist failed", zap.Error(err))
	}
	err = singleton.Query.DeleteEmailWhitelist(c, email)
	if err != nil {
		singleton.Logger.Error("remove email from whitelist table failed", zap.Error(err))
	}
	return
}
