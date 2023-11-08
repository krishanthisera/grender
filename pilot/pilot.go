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
		if err != nil {
			c.String(http.StatusInternalServerError, "Error rendering URL: %v", err)
			return
		}

		c.Data(http.StatusOK, "text/html", []byte(renderedHTML))
	})

	router.Run(":8080")
}
