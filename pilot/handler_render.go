package pilot

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krishanthisera/grender/backend"
)

// Rendering handler
func (C *Config) renderHandler(ctx *gin.Context) {
	url := ctx.Param("url")
	fmt.Println(url)

	var be backend.Backend
	var err error

	// Check S3 backend configuration
	if s3 := C.Backend.S3; s3.BucketName != "" {
		be, err = createBackendFromConfig(s3)
	} else if fs := C.Backend.FileSystem; fs.BaseDir != "" {
		// Check FileSystem backend configuration
		be, err = createBackendFromConfig(fs)
	} else {
		// No backend found
		fmt.Println("No backend found")
	}

	if err != nil {
		fmt.Println(err)
		ctx.String(http.StatusInternalServerError, "Error creating backend: %v", err)
		return
	}
	rac := renderAndCacheConfig{backend: &be, render: &C.RenderingConfig}
	renderedHTML, err := rac.RenderAndCache(url)

	// Page is rendered successfully
	if renderedHTML != nil {
		// Setting required headers
		C.addResponseHeaders(ctx)
		ctx.Data(http.StatusOK, "text/html", []byte(renderedHTML))

		// Page cannot be cached
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Error caching URL: %v", err)
			return
		}
		return
	}

	// Page cannot be rendered
	ctx.String(http.StatusInternalServerError, "Error rendering URL: %v", err)
}
