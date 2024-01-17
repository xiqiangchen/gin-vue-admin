package config

type Kafka struct {
	Brokers  []string  `json:"brokers" yaml:"brokers" mapstructure:"brokers"`
	Producer *Producer `json:"producer" yaml:"producer" mapstructure:"producer"`
	Consumer *Consumer `json:"consumer" yaml:"consumer" mapstructure:"consumer"`
}

type Producer struct {
	Topic    string `json:"topic" yaml:"topic" mapstructure:"topic"`
	Parallel int    `json:"parallel" yaml:"parallel" mapstructure:"parallel"`
}

type Consumer struct {
	Group    string   `json:"group" yaml:"group" mapstructure:"group"`
	Parallel int      `json:"parallel" yaml:"parallel" mapstructure:"parallel"`
	Topics   []string `json:"topics" yaml:"topics" mapstructure:"topics"`
	Oldest   bool     `json:"oldest" yaml:"oldest" mapstructure:"oldest"`
}
