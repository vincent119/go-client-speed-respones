package route

import (
	"github.com/gin-gonic/gin"
	//"github.com/vincent119/go-client-speed-respones/handle/web"
	//"github.com/vincent119/go-client-speed-respones/loggin"
	log4 "github.com/jeanphorn/log4go"
	"github.com/vincent119/go-client-speed-respones/config"
	tk "github.com/vincent119/go-client-speed-respones/handle/token"
	"github.com/vincent119/go-client-speed-respones/model"
	// "github.com/gomodule/redigo/redis"
	rds "github.com/vincent119/go-client-speed-respones/handle/rdsub"
	"fmt"
	"strings"
	"time"
)

func HandlGenToken (c *gin.Context){
	st := model.GenTokenString{}
	if tk.CheckHttpToken(c) == false {
	c.Abort()
	return
}
	err := c.BindJSON(&st)
	if err != nil {
	return
 }
	if len(st.ClinetIP) == 0 {
		fmt.Println("Client ip is empty...")
   	ipFor := c.GetHeader("x-forwarded-for") //ey = true
	 	fmt.Println("ipFor : " ,ipFor)
 }
  tokenSrring := config.GetServerSalt()+":"+st.Openid+":"+st.Uid+":"+ st.ClinetIP
  md5Value := tk.GenMd5(tokenSrring)
	fmt.Println("tokenSrring : " ,tokenSrring)
  fmt.Println("md5 : " ,md5Value)
	fmt.Println("256 : " ,tk.GenSha256(md5Value))
  rds.Set("11111","111123",1000)

}

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
