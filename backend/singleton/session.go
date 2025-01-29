package singleton

type Session struct {
	Key                     string `mapstructure:"key" json:"key" yaml:"key"`
	CookieName              string `mapstructure:"cookie-name" json:"cookie_name" yaml:"cookie-name"`
	ExpireTime              int    `mapstructure:"expire-time" json:"expire_time" yaml:"expire-time"`
	RefreshBeforeExpireTime int64  `mapstructure:"refresh-before-expire-time" json:"refresh_before_expire_time" yaml:"refresh-before-expire-time"`
	HttpOnly                bool   `mapstructure:"http-only" json:"http_only" yaml:"http-only"`
	Secure                  bool   `mapstructure:"secure" json:"secure" yaml:"secure"`
}
