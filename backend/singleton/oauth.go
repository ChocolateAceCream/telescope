package singleton

type OAuth struct {
	Google GoogleOAuthConfig `mapstructure:"google" json:"google" yaml:"google"`
}

type GoogleOAuthConfig struct {
	ClientID        string `mapstructure:"client-id" json:"client_id" yaml:"client-id"`
	ClientSecret    string `mapstructure:"client-secret" json:"client_secret" yaml:"client-secret"`
	RedirectURL     string `mapstructure:"redirect-url" json:"redirect_url" yaml:"redirect-url"`
	FrontendBaseUrl string `mapstructure:"frontend-base-url" json:"frontend_base_url" yaml:"frontend-base-url"`
}
