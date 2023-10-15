package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary SayHello
// @Tags auth
// @Description say hello
// @ID say-hello
// @Produce  json
// @Router /hello [get]
func (h *Handler) sayHello(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"hello": "world!",
	})
}
