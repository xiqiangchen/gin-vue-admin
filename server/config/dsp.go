package config

type Dsp struct {
	Env          string `mapstructure:"env" json:"env" yaml:"env"`                               // 环境值
	RouterPrefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"` // 路由前缀
	UseRedis     bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`             // 使用redis
	UseMongo     bool   `mapstructure:"use-mongo" json:"use-mongo" yaml:"use-mongo"`             // 使用mongo
	UseKafka     bool   `mapstructure:"use-kafka" json:"use-kafka" yaml:"use-kafka"`             // 使用kafka
	Domain       string `mapstructure:"domain" json:"domain" yaml:"domain"`
	Bid          Bid    `mapstructure:"bid" json:"bid" yaml:"bid"`
	Track        Track  `mapstructure:"track" json:"track" yaml:"track"`
	Ipv4Path     string `mapstructure:"ipv4-path" json:"ipv4-path" yaml:"ipv4-path"`
	Ipv6Path     string `mapstructure:"ipv6-path" json:"ipv6-path" yaml:"ipv6-path"`
}

type Bid struct {
	Port  int    `mapstructure:"port" json:"port" yaml:"port"` // 端口值
	Uri   string `mapstructure:"uri" json:"uri" yaml:"uri"`
	Topic string `json:"topic" yaml:"topic" mapstructure:"topic"`
}
type Track struct {
	Port       int        `mapstructure:"port" json:"port" yaml:"port"` // 端口值
	Impression Impression `mapstructure:"impression" json:"impression" yaml:"impression"`
	Click      Click      `mapstructure:"click" json:"click" yaml:"click"`
	Landing    Landing    `mapstructure:"landing" json:"landing" yaml:"landing"`
}

type Impression struct {
	Uri   string `mapstructure:"uri" json:"uri" yaml:"uri"`
	Topic string `json:"topic" yaml:"topic" mapstructure:"topic"`
}

type Click struct {
	Uri   string `mapstructure:"uri" json:"uri" yaml:"uri"`
	Topic string `json:"topic" yaml:"topic" mapstructure:"topic"`
}

type Landing struct {
	Uri   string `mapstructure:"uri" json:"uri" yaml:"uri"`
	Topic string `json:"topic" yaml:"topic" mapstructure:"topic"`
}
