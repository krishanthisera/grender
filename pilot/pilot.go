package pilot

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Grender() {
	router := gin.Default()
	router.GET("/render/*url", func(c *gin.Context) {
		url := c.Param("url")
		fmt.Println(url)
		renderedHTML, err := GetPages(url)

		// Page is rendered successfully
		if renderedHTML != nil {
			c.Data(http.StatusOK, "text/html", []byte(renderedHTML))
			// Page cannot be cached
			if err != nil {
				c.String(http.StatusInternalServerError, "Error caching URL: %v", err)
				return
			}
			return
		} else {
			// Page cannot be rendered
			c.String(http.StatusInternalServerError, "Error rendering URL: %v", err)
			return
		}

	})

	router.Run(":8080")
}
