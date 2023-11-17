package pilot

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (C *Config) Grender() {

	router := gin.Default()

	// Rendering requests
	router.GET("/render/*url", C.renderHandler)
	router.HEAD("/render/*url", C.renderHandler)

	router.Run(fmt.Sprintf(":%s", C.Server.Port))
}
