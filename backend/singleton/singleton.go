package singleton

import (
	"embed"

	sqlc "github.com/ChocolateAceCream/telescope/backend/db/sqlc"
	"github.com/redis/go-redis/v9"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	FileSystem  embed.FS
	Viper       *viper.Viper
	Logger      *zap.Logger
	Config      ServerConfig
	Redis       *redis.Client
	DB          *sqlc.Queries
	AWS         *AWSClient
	Translation map[string]map[string]string
)
