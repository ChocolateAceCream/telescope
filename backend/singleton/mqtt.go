package singleton

type MqttConfig struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	TlsPort  string `mapstructure:"tls-Port" json:"tls_port" yaml:"tls-Port"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	ClientId string `mapstructure:"client-id" json:"client_id" yaml:"client-id"`
}
