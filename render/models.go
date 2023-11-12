package render

type Config struct {
	PageWaitTime      float32 `yaml:"pageWaitTime"` // Seconds
	PageWailCondition string  `yaml:"pageWaitCondition"`
}
