package pilot

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (C *Config) Grender() {

	router := gin.Default()

	if C.Modes.Rendering {
		// Rendering requests
		router.GET("/render/*url", C.renderHandler)
		router.HEAD("/render/*url", C.renderHandler)
	}

	// Invalidation requests
	router.POST("/recache", C.invalidateHandler)

	if C.Modes.Recaching {
		go C.Invalidate.AMQP.Consumer(C.Backend.FileSystem)
	}

	// Start the server
	router.Run(fmt.Sprintf(":%s", C.Server.Port))

}
