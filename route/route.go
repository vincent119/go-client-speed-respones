package route
import (
	"net/http"
  "github.com/gin-gonic/gin"
	"github.com/vincent119/go-client-speed-respones/handle/web"
	//"github.com/vincent119/go-client-speed-respones/loggin"
	//"log"
	log4 "github.com/jeanphorn/log4go"
	 //"github.com/vincent119/go-client-speed-respones/handle/token"
)

func InitRouter() *gin.Engine {
	log4.LoadConfiguration("logging.json")
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
	api1.GET("/", web.HandleGet)
	// ping check
	api1.POST("/scheck", web.HandlePingCheck)
	// DNS check
	api1.POST("/dscheck", web.HandleDnsCheck)
	// client connect check
	api1.POST("/conncheck", web.HandleConnCheck)
	api1.GET("/healthcheck", web.HandleHealthCheck)


	return r
}