package main

import (
	"flag"
	"fmt"
	"go-client-speed-respones/config"
	"go-client-speed-respones/loggin"
	"github.com/gin-gonic/gin"
)

//var Conf =&Config{}


func main(){
  environment := flag.String("e", "dev", "")
  config.Init(*environment)
  Port := config.GetServerPort()
  ServerPort :=":"+Port
  ServerLog := config.GetServerLogPath()
  fmt.Println(ServerLog)
  //logging.InitializeLogging(ServerLog)
  //log.Println("wwwwwww")
  //log.Fatalf("What Happened??")
  loggin.MakeLogger(ServerLog,true)

  Routes := gin.Default()
  Routes.SetTrustedProxies([]string{"172.16.99.200"})
  Routes.Any("/",HandleGet)
  Routes.GET("/clinetrsp",HandleClinetResponse)
	/*Routes.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
				"message": "pong","aaaa":header_host,
		})
	})*/
  Routes.Run(ServerPort)
}
func HandleClinetResponse(c *gin.Context){
  for k,v :=range c.Request.Header {
		fmt.Println(k,v)
	}
}

func HandleGet(c *gin.Context) {
	c.JSON(200,gin.H{
		"receive":"65535",
	})
}

