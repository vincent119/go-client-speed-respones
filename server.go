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
  "strings"
  log4 "github.com/jeanphorn/log4go"
)


func main(){
  //environment := flag.String("e", "dev", "")
  config.Init()
  Port := config.GetServerPort()
  ServerPort :=":"+Port
  log4.LoadConfiguration("logging.json")

  Routes := gin.Default()
  // Server log init
  Routes.Use(loggin.LoggerToFile(config.GetServerLogFile()))
  Routes.SetTrustedProxies([]string{"172.16.99.200"})
  Routes.GET("/",HandleGet)
  // ping check
  Routes.POST("/pcheck",HandlePingCheck)
  // DNS check
  Routes.POST("/dscheck",HandleDnsCheck)
  //
  Routes.POST("/conncheck",HandleConnCheck)
  Routes.GET("/healthcheck",HandleHealthCheck)
  Routes.Run(ServerPort)
}

func CheckHttpToken(c *gin.Context) bool{
  TokenValues := c.GetHeader("token")
  if config.GetServerToken() != TokenValues {
    c.JSON(401,gin.H{
      "Status": "401",
    })
    return false
  } else {
    return true
  }
}

func HandleConnCheck(c *gin.Context){
  st := model.ClientConnStatus{}
  //token = c.Request.Header["Token"]
  if CheckHttpToken(c) == false {
    c.Abort()
    return
  }
  err := c.BindJSON(&st)
  if err != nil {
    return
  }
  log4.LOGGER("connectcheck").Info(strings.Replace(fmt.Sprintf("%#v", st), ", ",",", -1))
  c.JSON(200,gin.H{
    "Status":"OK", "recv_time": fmt.Sprint(time.Now().Format("2006/1/2 15:04:05.999")),
  })
}

func HandleDnsCheck(c *gin.Context){
  st := model.ClientDnsStatus{}
  if CheckHttpToken(c) == false {
    c.Abort()
    return
  }
  err := c.BindJSON(&st)
  if err != nil {
    return
  }
  fmt.Print(st)
  log4.LOGGER("dnscheck").Info(strings.Replace(fmt.Sprintf("%#v", st), ", ",",", -1))
  c.JSON(200,gin.H{
    "Status":"OK", "recv_time": fmt.Sprint(time.Now().Format("2006/1/2 15:04:05.999")),
  })
}

func HandlePingCheck(c *gin.Context){
  md := model.ClientPingStatus{}
  if CheckHttpToken(c) == false {
    c.Abort()
    return
  }
  err := c.BindJSON(&md)
  if err != nil {
    return
  }

  log4.LOGGER("pingcheck").Info(strings.Replace(fmt.Sprintf("%#v", md), ", ",",", -1))
  c.JSON(200,gin.H{
    "Status":"OK", "recv_time": fmt.Sprint(time.Now().Format("2006/1/2 15:04:05.999")),
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

