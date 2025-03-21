package singleton

type Lock struct {
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Timeout      int64  `mapstructure:"timeout" json:"timeout" yaml:"timeout"`
	TimeInterval int64  `mapstructure:"time-interval" json:"time_interval" yaml:"time-interval"`
	HttpOnly     bool   `mapstructure:"http-only" json:"http_only" yaml:"http-only"`
	Secure       bool   `mapstructure:"secure" json:"secure" yaml:"secure"`
}
