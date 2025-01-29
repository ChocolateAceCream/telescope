package singleton

type Workers struct {
	Resizer Resizer `mapstructure:"resizer" json:"resizer" yaml:"resizer"`
}

type Resizer struct {
	Count     int `mapstructure:"count" json:"count" yaml:"count"`
	QueueSize int `mapstructure:"queue-size" json:"queue_size" yaml:"queue-size"`
}
