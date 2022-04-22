package route

import (
	"github.com/gin-gonic/gin"
	//"github.com/vincent119/go-client-speed-respones/handle/web"
	//"github.com/vincent119/go-client-speed-respones/loggin"
	"fmt"
	"strings"
	"time"

	//"github.com/google/uuid"
	log4 "github.com/jeanphorn/log4go"
	"github.com/vincent119/go-client-speed-respones/config"
	co "github.com/vincent119/go-client-speed-respones/handle/crypto"
	"github.com/vincent119/go-client-speed-respones/handle/ctime"
	rds "github.com/vincent119/go-client-speed-respones/handle/rdsub"
	"github.com/vincent119/go-client-speed-respones/model"

	//"encoding/json"
	"os"
	"reflect"
	"io/ioutil"
)


func HandleGenToken(c *gin.Context) {
	st := model.GenTokenString{}
	//x, _ := ioutil.ReadAll(c.Request.Body)
  //fmt.Printf("%s\n", string(x))
	err:= c.ShouldBind(&st)
	//err := c.BindJSON(&st)
	//err:= c.ShouldBind(&st)
	if err != nil {
		fmt.Println("json data error")
		return
	}

	fmt.Println(strings.Replace(fmt.Sprintf("%+v", st), ", ", ", ", -1))
	if len(st.ClinetIP) == 0 {
		fmt.Println("Client ip is empty...")
		//ipFor := c.GetHeader("x-forwarded-for") //ey = true
		//fmt.Println(c.ClientIP())
		//fmt.Println("ipFor : ", ipFor)
		st.ClinetIP = c.ClientIP()
	}
	// tokenString = openid + Ukey + Client IP
	//tokenString := config.GetServerSalt() + "." + st.Openid + "." + st.Ukey + ":" + st.ClinetIP
	//uuid := uuid.New()
	//fmt.Println(uuid)
  tokenString :=  config.GetServerUkey()+":"+st.Openid + ":" + st.Uid + ":" + st.ClinetIP+":"+config.GetServerVtoken()
	//tokenString := st.Openid+":"+st.Ukey+":"+ st.ClinetIP
	md5Value := co.GenMd5(tokenString)
	md5ValueS := md5Value + ":" + config.GetServerSalt()
	sha256Value := co.GenSha256(md5ValueS)
	rds.Set(md5Value, sha256Value, config.RedisTtl())
	log4.LOGGER("genkon").Info(strings.Replace(fmt.Sprintf("%#v", st), ", ", ",", -1))
	//log4.LOGGER("gentok").Info("value: %s", sha256Value)
	c.Header("x-key", md5Value)
	//c.Header("uyccc","5555333555")
	//c.Header("Access-Control-Allow-Origin","*")
	c.JSON(200, gin.H{
		"Status": "OK", "recv_time": fmt.Sprint(time.Now().Format("2006-01-02T15:04:05")),
	})
}

func HandleConnCheck(c *gin.Context) {
	st := model.ClientConnStatus{}
	err:= c.ShouldBind(&st)
	//err := c.BindJSON(&st)
	if err != nil {
		return
	}
	st.Category =  reflect.TypeOf(st).String()
	st.Host, _ = os.Hostname()
	t := st.TimeStamp
	//fmt.Println(ctime.FormatIso88601(t))
	st.TimeStamp = ctime.FormatIso88601(t)
	fmt.Println(strings.Replace(fmt.Sprintf("%+v", st), ", ", ", ", -1))
	log4.LOGGER("connectcheck").Info(strings.Replace(fmt.Sprintf("%+v", st), ", ", ",", -1))
	c.JSON(200, gin.H{
		"Status": "OK", "recv_time": fmt.Sprint(time.Now().Format("2006-01-02T15:04:05")),
	})
}

// @Success 200 {string} string
// @Router /dsheck [post]
func HandleDnsCheck(c *gin.Context) {
	st := model.ClientDnsStatus{}
	err:= c.ShouldBind(&st)
	//err := c.BindJSON(&st)
	if err != nil {
		return
	}
	st.Category =  reflect.TypeOf(st).String()
	st.Host,_ = os.Hostname()
	t := st.TimeStamp
	//fmt.Println(ctime.FormatIso88601(t))
	st.TimeStamp = ctime.FormatIso88601(t)
	//fmt.Println(strings.Replace(fmt.Sprintf("%+v", st), ", ", ", ", -1))
	log4.LOGGER("dnscheck").Info(strings.Replace(fmt.Sprintf("%+v", st), ", ", ",", -1))
	c.JSON(200, gin.H{
		"Status": "OK", "recv_time": fmt.Sprint(time.Now().Format("2006-01-02T15:04:05")),
	})
}

// @summary ping Status Check
// @Success 200 {string} string
// @Router /scheck [post]
func HandlePingCheck(c *gin.Context) {
  md := model.ClientPingStatus{}
	//md := model.PingStatus{}
	err:= c.ShouldBind(&md)
	//err := c.BindJSON(&md)
	if err != nil {
		return
	}
	//fmt.Println(reflect.TypeOf(md).String())
	//md.pk = make(map[string]string)
	//md.pk["cat"] = "55688"
	//fmt.Println(reflect.TypeOf(md).String())
	//aa := reflect.TypeOf(md).String()+string(jsonE)
	//md.Message = reflect.TypeOf(md)
	md.Category = reflect.TypeOf(md).String()
	md.Host,_  = os.Hostname()
	t := md.TimeStamp
	fmt.Println(ctime.FormatIso88601(t))
	md.TimeStamp = ctime.FormatIso88601(t)
	fmt.Println(strings.Replace(fmt.Sprintf("%+v", md), ", ", ", ", -1))
	log4.LOGGER("pingcheck").Info(strings.Replace(fmt.Sprintf("%+v", md), ", ", ", ", -1))
	c.JSON(200, gin.H{
		"Status": "OK", "recv_time": fmt.Sprint(time.Now().Format("2006-01-02T15:04:05")),
	})
}

// @summary Check health of api service
// @Success 200 {string} string
// @Router /healthcheck [get]
func HandleHealthCheck(c *gin.Context) {
	//token = c.Request.Header["Token"]
	c.JSON(200, gin.H{
		"Health Status": "OK", "recv_time": fmt.Sprint(time.Now().Format("2006-01-02T15:04:05")),
	})
}

// @Success 200 {string} string
// @Router / [get]
// get root path 
func HandleGet(c *gin.Context) {
	//Routes.Use(loggin.LoggerToFile())
	c.JSON(200, gin.H{
		"receive": "65535", "time": fmt.Sprint(time.Now().Format("2006-01-02T15:04:05")),
	})
}

func HttpHeaderGet(c *gin.Context) {
	for k,v := range c.Request.Header {
     fmt.Println(k,v)
	}
}

func HttpBodyGet(c *gin.Context) {
	x, _ := ioutil.ReadAll(c.Request.Body)
  fmt.Printf("%s\n", string(x))
}
