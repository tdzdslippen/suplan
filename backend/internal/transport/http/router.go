package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() http.Handler {
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	return r
}
