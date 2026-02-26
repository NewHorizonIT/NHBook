package logger

type ConfigLogger struct {
	// Zap logger configuration
	Level         string `mapstructure:"level" json:"level" yaml:"level"`
	Format        string `mapstructure:"format" json:"format" yaml:"format"`
	EnableConsole bool   `mapstructure:"enable_console" json:"enable_console" yaml:"enable_console"`
	AppLog        string `mapstructure:"app_log" json:"app_log" yaml:"app_log"`
	RequestLog    string `mapstructure:"request_log" json:"request_log" yaml:"request_log"`
	ErrorLog      string `mapstructure:"error_log" json:"error_log" yaml:"error_log"`

	// Lumberjack configuration
	Filename   string `mapstructure:"filename" json:"filename" yaml:"filename"`
	MaxSize    int    `mapstructure:"max_size" json:"max_size" yaml:"max_size"`
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"`
	MaxAge     int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`
	Compress   bool   `mapstructure:"compress" json:"compress" yaml:"compress"`
}
