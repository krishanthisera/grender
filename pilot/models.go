package pilot

import (
	"github.com/krishanthisera/grender/backend"
	"github.com/krishanthisera/grender/render"
)

// Config struct to represent the overall YAML configuration
type Config struct {
	Version         string        `yaml:"version" required:"true"`
	RenderingConfig render.Config `yaml:"renderingConfig" required:"true"`
	Server          struct {
		Port            string `yaml:"port" required:"true"`
		ResponseHeaders []struct {
			Name  string `yaml:"name"`
			Value string `yaml:"value"`
		} `yaml:"responseHeaders"`
	} `yaml:"server" required:"true"`
	Backend struct {
		S3         backend.S3         `yaml:"s3"`
		FileSystem backend.FileSystem `yaml:"fileSystem"`
	} `yaml:"backend" required:"true"`
}

type renderAndCacheConfig struct {
	backend *backend.Backend
	render  *render.Config
}
