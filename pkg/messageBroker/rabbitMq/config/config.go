package rabbitMqconfig

type RabbitMQConfig struct {
	PrefetchCount  int    `yaml:"prefetch-count"`
	GoroutineLimit uint64 `yaml:"go-routine-limit"`
	QueueName      string `yaml:"queue-name"`
	ConnectionURL  string `yaml:"connection-url"`
}
