package main

import (
	//"net/http"
	"github.com/gin-gonic/gin"
	"github.com/vincent119/go-client-speed-respones/config"
	tk "github.com/vincent119/go-client-speed-respones/handle/token"
	"github.com/vincent119/go-client-speed-respones/loggin"
	rt "github.com/vincent119/go-client-speed-respones/route"

	//hp "github.com/vincent119/go-client-speed-respones/handle/http"
	//"fmt"
	//"strings"
	//"time"
	//"go-client-speed-respones/loggin"
	//"github.com/vincent119/go-client-speed-respones/model"
	log4 "github.com/jeanphorn/log4go"
	"github.com/vincent119/go-client-speed-respones/handle/rdsub"
	//"github.com/vincent119/go-client-speed-respones/middleware/rdsub"
)

func init() {
	config.Init()
	rdsub.Setup()
}

// @title Gin
// @version 1.0
// @description Gin API
// @contact.name Vincent Yu
// @host localhost:8080
// @schemes http

func main() {
	//config.Init()
	//router := route.InitRouter()
	//model.RedisConnection()

	//model.RedisInit()
	//model.RdbGet()
	//environment := flag.String("e", "dev", "")
	Port := config.GetServerPort()
	ServerPort := ":" + Port
	log4.LoadConfiguration("logging.json")

	Routes := gin.Default()
	Routes.Use(gin.Logger(), gin.Recovery())
	//Server log init
	Routes.Use(loggin.LoggerToFile(config.GetServerLogFile()))
	Routes.SetTrustedProxies([]string{"172.16.99.200"})
	Routes.GET("/", rt.HandleGet)
	// ping check
	Routes.POST("/scheck", tk.CheckHttpToken,tk.CheckHttpXkey,tk.CheckRdbXkey, rt.HandlePingCheck)
	// DNS check
	Routes.POST("/dscheck", rt.HandleDnsCheck)
	// client connect check
	Routes.POST("/conncheck", tk.CheckHttpToken, rt.HandleConnCheck)
	Routes.GET("/healthcheck", rt.HandleHealthCheck)
	Routes.POST("/gentok", tk.CheckHttpToken, rt.HandlGenToken)

	Routes.Run(ServerPort)
}

// func CheckHttpToken(c *gin.Context) bool {
// 	TokenValues := c.GetHeader("utoken")
// 	if config.GetServerUkey() != TokenValues {
// 		c.JSON(401, gin.H{
// 			"Status": "401",
// 		})
// 		return false
// 	} else {
// 		return true
// 	}
// }

// @summary connect check fir Client
// @Success 200 {string} string
// @Router /conncheck [post]
// @produce application/json;charset=utf-8
// @param clientIp path string true "10.10.1.1"
// @param domain path string true "www.abc.com"
// @param time path string true "2022/02/18 12:25:48.32"
// @param status path string true "can not connect"
/*
func HandleConnCheck(c *gin.Context) {
	st := model.ClientConnStatus{}
	//token = c.Request.Header["Token"]
	if tk.CheckHttpToken(c) == false {
		c.Abort()
		return
	}
	err := c.BindJSON(&st)
	if err != nil {
		return
	}
	log4.LOGGER("connectcheck").Info(strings.Replace(fmt.Sprintf("%#v", st), ", ", ",", -1))
	c.JSON(200, gin.H{
		"Status": "OK", "recv_time": fmt.Sprint(time.Now().Format("2006/1/2 15:04:05.999")),
	})
}

// @Success 200 {string} string
// @Router /dsheck [post]
func HandleDnsCheck(c *gin.Context) {
	st := model.ClientDnsStatus{}
	if tk.CheckHttpToken(c) == false {
		c.Abort()
		return
	}
	err := c.BindJSON(&st)
	if err != nil {
		return
	}
	fmt.Print(st)
	log4.LOGGER("dnscheck").Info(strings.Replace(fmt.Sprintf("%#v", st), ", ", ",", -1))
	c.JSON(200, gin.H{
		"Status": "OK", "recv_time": fmt.Sprint(time.Now().Format("2006/1/2 15:04:05.999")),
	})
}

// @summary ping Status Check
// @Success 200 {string} string
// @Router /pcheck [post]
func HandlePingCheck(c *gin.Context) {
	md := model.ClientPingStatus{}
	if tk.CheckHttpToken(c) == false {
		c.Abort()
		return
	}
	err := c.BindJSON(&md)
	if err != nil {
		return
	}

	log4.LOGGER("pingcheck").Info(strings.Replace(fmt.Sprintf("%#v", md), ", ", ",", -1))
	c.JSON(200, gin.H{
		"Status": "OK", "recv_time": fmt.Sprint(time.Now().Format("2006/1/2 15:04:05.999")),
	})
}

// @summary Check health of api service
// @Success 200 {string} string
// @Router /healthcheck [get]
func HandleHealthCheck(c *gin.Context) {
	//token = c.Request.Header["Token"]
	c.JSON(200, gin.H{
		"Health Status": "OK", "recv_time": fmt.Sprint(time.Now().Format("2006/1/2 15:04:05.999")),
	})
}

// @Success 200 {string} string
// @Router / [get]
func HandleGet(c *gin.Context) {
	//Routes.Use(loggin.LoggerToFile())
	c.JSON(200, gin.H{
		"receive": "65535", "time": fmt.Sprint(time.Now().Format("2006/1/2 15:04:05.999")),
	})
}
*/
