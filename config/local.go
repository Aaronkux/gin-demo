package config

type Local struct {
	Path   string `mapstructure:"path" json:"path" yaml:"path"`
	Avatar string `mapstructure:"avatar" json:"avatar" yaml:"avatar"`
}
