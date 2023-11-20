package invalidate

type Queue interface {
	Push(payload *[]string) error
	Consume()
}

type AMQP struct {
	URI       string `yaml:"uri"`
	QueueName string `yaml:"queue" default:"invalidate"`
	TimeOut   int    `yaml:"timeout"`
	Exchange  string `yaml:"exchange" default:""`
}

type SQS struct{}
