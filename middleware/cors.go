package middleware

import (
	//"fmt"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		  //fmt.Println(c.Request.Header["Origin"][0])
		  if value ,_ := c.Request.Header["Origin"]; len(value) > 0 {
				c.Writer.Header().Set("Access-Control-Allow-Origin", value[0])
			} else {
				c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			}
			//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			//c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With,utoken,x-key")
			c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,utoken,x-key")	
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
			// for Google Chrome
			c.Writer.Header().Set("Referrer-Policy ", "origin")
		
			if c.Request.Method == "OPTIONS" {
				  c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header["Origin"][0])
					c.AbortWithStatus(204)
					return
			}
			c.Next()
	}
}