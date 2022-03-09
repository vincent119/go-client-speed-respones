package token

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vincent119/go-client-speed-respones/config"
)

func CheckHttpToken(c *gin.Context) bool {
	TokenValues := c.GetHeader("utoken")
	if config.GetServerUkey() != TokenValues {
		c.JSON(401, gin.H{
			"Status": "401",
		})
		return false
	} else {
		return true
	}
}