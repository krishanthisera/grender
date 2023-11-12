package pilot

import (
	"github.com/krishanthisera/grender/backend"
	"github.com/krishanthisera/grender/render"
)

// Config struct to represent the overall YAML configuration
type Config struct {
	Version         string        `yaml:"version"`
	RenderingConfig render.Config `yaml:"renderingConfig"`
	Server          struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Backend struct {
		S3         backend.S3         `yaml:"s3"`
		FileSystem backend.FileSystem `yaml:"fileSystem"`
	} `yaml:"backend"`
}

type renderAndCacheConfig struct {
	backend *backend.Backend
	render  *render.Config
}
