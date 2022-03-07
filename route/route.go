package route
import (
	"net/http"
  "github.com/gin-gonic/gin"
	 "github.com/vincent119/go-client-speed-respones/handle/webaccess"
	 "github.com/vincent119/go-client-speed-respones/handle/token"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	// 避免404
	r.NoRoute(func(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", gin.H{})
	})
	r.NoMethod(func(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", gin.H{})
	})

	api1 := r.Group("")
	api1.GET("/", acc.HandleGet)
	// ping check
	api1.POST("/scheck", acc.HandlePingCheck)
	// DNS check
	api1.POST("/dscheck", acc.HandleDnsCheck)
	// client connect check
	api1.POST("/conncheck", acc.HandleConnCheck)
	api1.GET("/healthcheck", acc.HandleHealthCheck)


	return r
}