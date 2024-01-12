package config

type Dsp struct {
	Env          string `mapstructure:"env" json:"env" yaml:"env"`                               // 环境值
	RouterPrefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"` // 路由前缀
	Addr         int    `mapstructure:"addr" json:"addr" yaml:"addr"`                            // 端口值
	UseRedis     bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`             // 使用redis
	UseMongo     bool   `mapstructure:"use-mongo" json:"use-mongo" yaml:"use-mongo"`             // 使用mongo
	UseKafka     bool   `mapstructure:"use-kafka" json:"use-kafka" yaml:"use-kafka"`             // 使用kafka
}
