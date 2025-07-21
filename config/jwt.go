package config

// Jwt token
type Jwt struct {
	Key        string `yaml:"key"`
	Exp        int64  `yaml:"exp"`
	RefreshExp int64  `yaml:"refresh-exp"`
}
