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
	"github.com/vincent119/go-client-speed-respones/middleware"
	"github.com/gin-contrib/gzip"
	//cors  "github.com/gin-contrib/cors"
	//"time"
//	"github.com/dvwright/xss-mw"
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
func setupRoute() *gin.Engine{
  //var xssMdlwr xss.XssMw
	// Disable Print to console color 
	gin.DisableConsoleColor()
	//config.Init()
	//router := route.InitRouter()
	//model.RedisConnection()

	//model.RedisInit()
	//model.RdbGet()
	//environment := flag.String("e", "dev", "")
	
	log4.LoadConfiguration("logging.json")

	routes := gin.Default()
	routes.Use(gin.Logger(), 
	//xssMdlwr.RemoveXss(),
	gin.Recovery(),
	middleware.CORSMiddleware(),
	gzip.Gzip(gzip.DefaultCompression),
  )
	//routes.Ues(middleware.CORSMiddleware())
	//corsConfig := cors.DefaultConfig()
	//corsConfig.AllowOrigins = []string{"*"}
	//corsConfig.AllowAllOrigins = true
	// corsConfig.AllowHeaders = []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "Origin", "Cache-Control", "X-Requested-With","Authorization","Connection","Host","Access-Control-Request-Method","x-key","utoken"}
	// corsConfig.AllowMethods = []string{"POST","OPTIONS","GET"}
	// corsConfig.ExposeHeaders = []string{"Content-Length,x-key,utoken"}
	// //corsConfig.AllowCredentials = true
	// corsConfig.MaxAge = 12 * time.Hour
	
	//routes.Use(cors.New(corsConfig))
	/* routes.Use(cors.New(cors.config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST"},
		AllowHeaders:     []string{"Origin,Content-Type,x-key,utoken"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	 })) */
	//Server log init
	routes.Use(loggin.LoggerToFile(config.GetServerLogFile()))
	routes.SetTrustedProxies([]string{"172.16.99.200"})
	routes.GET("/", rt.HandleGet)
	// ping check
	routes.POST("/scheck", rt.HttpHeaderGet,tk.CheckHttpToken,tk.CheckHttpXkey,tk.CheckRdbXkey, rt.HandlePingCheck)
	// DNS check
	routes.POST("/dscheck", rt.HandleDnsCheck)
	// client connect check
	routes.POST("/conncheck", tk.CheckHttpToken,tk.CheckHttpXkey,tk.CheckRdbXkey,rt.HandleConnCheck)
	routes.GET("/healthcheck", rt.HandleHealthCheck)
	// Route get token 
	// HttpHeaderGet  get Http request header 
	routes.POST("/gentok", rt.HttpHeaderGet, tk.CheckHttpToken, rt.HandleGenToken)
	//routes.POST("/gentok", tk.CheckHttpToken, rt.HandleGenToken)
	return routes
}


func main() {
	Port := config.GetServerPort()
	ServerPort := ":" + Port
	r := setupRoute()
	r.Run(ServerPort)
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
