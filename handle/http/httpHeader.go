package http

import (
	"github.com/gin-gonic/gin"

)

func HeaderXkey(c *gin.Context) string {
	return c.GetHeader("x-key")
}
func HeaderUtoken(c *gin.Context) string {
	return c.GetHeader("utoken")
}
func HeaderUkey(c *gin.Context) string {
	return c.GetHeader("ukey")
}
