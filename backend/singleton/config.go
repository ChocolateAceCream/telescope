package singleton

type ServerConfig struct {
	Zap       Zap         `mapstructure:"zap" json:"zap" yaml:"zap"`
	Mysql     Mysql       `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis     RedisConfig `mapstructure:"redis" json:"redis" yaml:"redis"`
	Session   Session     `mapstructure:"session" json:"session" yaml:"session"`
	Lock      Lock        `mapstructure:"lock" json:"lock" yaml:"lock"`
	Limiter   Limiter     `mapstructure:"limiter" json:"limiter" yaml:"limiter"`
	Captcha   Captcha     `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Email     Email       `mapstructure:"email" json:"email" yaml:"email"`
	Signature Signature   `mapstructure:"signature" json:"signature" yaml:"signature"`
	Init      Init        `mapstructure:"init" json:"init" yaml:"init"`
	Local     Local       `mapstructure:"local" json:"local" yaml:"local"`
	Minio     Minio       `mapstructure:"minio" json:"minio" yaml:"minio"`
	Mqtt      MqttConfig  `mapstructure:"mqtt" json:"mqtt" yaml:"mqtt"`
	DB        DBConfig    `mapstructure:"db" json:"db" yaml:"db"`
	OAuth     OAuth       `mapstructure:"oauth" json:"oauth" yaml:"oauth"`
	Workers   Workers     `mapstructure:"workers" json:"workers" yaml:"workers"`
	AWS       AWSConfig   `mapstructure:"aws" json:"aws" yaml:"aws"`
}
