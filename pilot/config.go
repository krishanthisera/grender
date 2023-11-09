package pilot

import (
	"github.com/krishanthisera/grender/render"
)

type Pilot struct {
	RenderConfigs render.RenderingConfigs `yaml:"rendering_configs"`
}
