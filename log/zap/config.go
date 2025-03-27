package zap

type Config struct {
	Level  string       `mapstructure:"level"`
	File   FileConfig   `mapstructure:"file"`
	Stdout StdoutConfig `mapstructure:"stdout"`
	KV     KVConfig     `mapstructure:"kv"`
}

type FileConfig struct {
	Filename   string `mapstructure:"filename"`
	Maxsize    int64  `mapstructure:"maxsize"`
	Maxbackups int64  `mapstructure:"maxbackups"`
	Maxage     int64  `mapstructure:"maxage"`
	Compress   bool   `mapstructure:"compress"`
	Enable     bool   `mapstructure:"enable"`
}

type StdoutConfig struct {
	Enable bool `mapstructure:"enable"`
}

type KVConfig struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
	ID      string `mapstructure:"id"`
	Trace   bool   `mapstructure:"trace"`
}
