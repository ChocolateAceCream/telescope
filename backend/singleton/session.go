package singleton

type Session struct {
	CookieName              string `mapstructure:"cookie-name" json:"cookie_name" yaml:"cookie-name"`
	RefreshTokenExpireTime  int    `mapstructure:"refresh-token-expire-time" json:"refresh_token_expire_time" yaml:"refresh-token-expire-time"`
	ExpireTime              int    `mapstructure:"expire-time" json:"expire_time" yaml:"expire-time"`
	RefreshBeforeExpireTime int64  `mapstructure:"refresh-before-expire-time" json:"refresh_before_expire_time" yaml:"refresh-before-expire-time"`
	HttpOnly                bool   `mapstructure:"http-only" json:"http_only" yaml:"http-only"`
	Secure                  bool   `mapstructure:"secure" json:"secure" yaml:"secure"`
}
