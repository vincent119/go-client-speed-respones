package web

import (
	"github.com/gin-gonic/gin"
  "time"
	"fmt"
	"strings"
	log4 "github.com/jeanphorn/log4go"
	"github.com/vincent119/go-client-speed-respones/model"
	"github.com/vincent119/go-client-speed-respones/handle/token"
)

func HandleConnCheck(c *gin.Context) {
	st := model.ClientConnStatus{}
	//token = c.Request.Header["Token"]
	if token.CheckHttpToken(c) == false {
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
	if token.CheckHttpToken(c) == false {
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
	if token.CheckHttpToken(c) == false {
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
		"Status": "OK", "recv_time": fmt.Sprint(time.Now().Format("2006/1/2 15:04:05.999")),
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