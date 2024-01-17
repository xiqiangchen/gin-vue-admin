package config

type Dsp struct {
	Env          string `mapstructure:"env" json:"env" yaml:"env"`                               // 环境值
	RouterPrefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"` // 路由前缀
	UseRedis     bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`             // 使用redis
	UseMongo     bool   `mapstructure:"use-mongo" json:"use-mongo" yaml:"use-mongo"`             // 使用mongo
	UseKafka     bool   `mapstructure:"use-kafka" json:"use-kafka" yaml:"use-kafka"`             // 使用kafka
	Bid          Bid    `mapstructure:"bid" json:"bid" yaml:"bid"`
	Track        Track  `mapstructure:"track" json:"track" yaml:"track"`
}

type Bid struct {
	Port int    `mapstructure:"port" json:"port" yaml:"port"` // 端口值
	Uri  string `mapstructure:"uri" json:"uri" yaml:"uri"`
}
type Track struct {
	Port          int    `mapstructure:"port" json:"port" yaml:"port"` // 端口值
	ImpressionUri string `mapstructure:"impression-uri" json:"impression-uri" yaml:"impression-uri"`
	ClickUri      string `mapstructure:"click-uri" json:"click-uri" yaml:"click-uri"`
}
