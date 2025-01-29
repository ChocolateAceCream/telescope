package singleton

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`           // log prefix
	Format        string `mapstructure:"format" json:"format" yaml:"format"`           // output format
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`    // log storage dir
	FileName      string `mapstructure:"file-name" json:"file_name"  yaml:"file-name"` // log file name
	EncodeLevel   string `mapstructure:"encode-level" json:"encode_level" yaml:"encode-level"`
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktrace_key" yaml:"stacktrace-key"`

	MaxAge       int  `mapstructure:"max-age" json:"max_age" yaml:"max-age"`       // log live age
	ShowLine     bool `mapstructure:"show-line" json:"show_line" yaml:"show-line"` // show callee line
	LogInConsole bool `mapstructure:"log-in-console" json:"log_in_console" yaml:"log-in-console"`
}

// InteroperateLevel: return zapcore.Level based on passed in config Level value
// func (z *Zap) TransportLevel() zapcore.Level {
// 	z.Level = strings.ToLower(z.Level)
// 	switch z.Level {
// 	case "debug":
// 		return zapcore.DebugLevel
// 	case "info":
// 		return zapcore.InfoLevel
// 	case "warn":
// 		return zapcore.WarnLevel
// 	case "error":
// 		return zapcore.WarnLevel
// 	case "dpanic":
// 		return zapcore.DPanicLevel
// 	case "panic":
// 		return zapcore.PanicLevel
// 	case "fatal":
// 		return zapcore.FatalLevel
// 	default:
// 		return zapcore.DebugLevel
// 	}
// }
