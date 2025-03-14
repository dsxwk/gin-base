package config

type Env struct {
	Mode string `yaml:"mode"` // 运行环境模式 debug, test, release
	Port string `yaml:"port"` // 服务端口
}
