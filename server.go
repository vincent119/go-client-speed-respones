package main

import (
	//"flag"
	"fmt"
  "github.com/vincent119/go-client-speed-respones/config"
  "github.com/vincent119/go-client-speed-respones/loggin" 
   "github.com/vincent119/go-client-speed-respones/model" 
	"github.com/gin-gonic/gin"
  //"log"
  "time"

)




func main(){
  //environment := flag.String("e", "dev", "")
  config.Init()
  Port := config.GetServerPort()
  ServerPort :=":"+Port
  //ServerLog := config.GetServerLogPath()
  //loggin.MakeLogger(ServerLog,true)
  //fmt.Println(ServerLog)
  //log.Println(ServerPort)

  Routes := gin.Default()
  //logFileName := config.GetServerLogFile()
  Routes.Use(loggin.LoggerToFile(config.GetServerLogFile()))
  Routes.SetTrustedProxies([]string{"172.16.99.200"})
  Routes.GET("/",HandleGet)
  Routes.POST("/clinetrsp",HandleClinetResponse)
  Routes.GET("/healthcheck",HandleHealthCheck)
	/*Routes.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
				"message": "pong","aaaa":header_host,
		})
	})*/
  Routes.Run(ServerPort)
}
func HandleClinetResponse(c *gin.Context){
  //config.Init()
  //fmt.Println(config.GetUrl1LogName())
  //Routes := gin.Default()
  //Routes.Use(loggin.LoggerToFile(config.GetUrl1LogName()))
  //for k,v :=range c.Request.Header {
	//	fmt.Println(k,v)
	//}
  //md := make(map[string]interface{})
  md := model.UrlClinetrsp{}
  //fmt.Println(c.Request.Header)
  err := c.BindJSON(&md)
  if err != nil {
    return
  }
  fmt.Printf("%v\n" ,&md)
  c.JSON(200,gin.H{
    "Status":"OK", "recv_time": fmt.Sprint(time.Now().Format("2006/1/2 15:04:05.999")),
    "ip": md.ClinetIP,
  })
}

func HandleHealthCheck(c *gin.Context){
  //token = c.Request.Header["Token"]

  c.JSON(200,gin.H{
    "Status":"OK", "recv_time": fmt.Sprint(time.Now().Format("2006/1/2 15:04:05.999")),
  })
}

func HandleGet(c *gin.Context) {
  //Routes.Use(loggin.LoggerToFile())
	c.JSON(200,gin.H{
		"receive":"65535","time": fmt.Sprint(time.Now().Format("2006/1/2 15:04:05.999")),
  })
}

