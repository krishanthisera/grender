package render

type Config struct {
	PageWaitTime      float32 `yaml:"pageWaitTime"` // Seconds
	PageWailCondition string  `yaml:"pageWaitCondition"`
	RequestHeaders    []struct {
		Name  string `yaml:"name"`
		Value string `yaml:"value"`
	} `yaml:"requestHeaders"`
}
