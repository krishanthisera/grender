package pilot

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (C *Config) invalidateHandler(ctx *gin.Context) {

	var urls *[]string
	if err := ctx.BindJSON(&urls); err != nil {
		ctx.String(http.StatusInternalServerError, "Error parsing JSON: %v", err)
		return
	}

	if err := C.Invalidate.AMQP.Push(urls); err != nil {
		ctx.String(http.StatusInternalServerError, "Error pushing to AMQP: %v", err)
		return
	}

	ctx.IndentedJSON(http.StatusCreated, urls)
}
