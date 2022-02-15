package main

import (
	//"flag"
	"fmt"
  "github.com/vincent119/go-client-speed-respones/config"
  "github.com/vincent119/go-client-speed-respones/loggin" 
	"github.com/gin-gonic/gin"
  "log"
)

//var Conf =&Config{}


func main(){
  //environment := flag.String("e", "dev", "")
  config.Init()
  Port := config.GetServerPort()
  ServerPort :=":"+Port
  ServerLog := config.GetServerLogPath()
  loggin.MakeLogger(ServerLog,true)
  fmt.Println(ServerLog)
  //loggin.InitializeLogging(ServerLog)
  log.Println(ServerPort)
  //log.Fatalf("What Happened??")


  

  Routes := gin.Default()
  Routes.Use(loggin.LoggerToFile(config.GetServerLogPath()))
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

