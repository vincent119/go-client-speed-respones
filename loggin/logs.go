package loggin

import (
	"fmt"
	"github.com/vincent119/go-client-speed-respones/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

func Logger(ph string) *logrus.Logger {

	config.Init()
	logFilePath := config.GetServerLogPath()
	logFileName := ph
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
	}

	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}

	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	logger := logrus.New()
	logger.Out = src
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-1-2 15:04:05.9999",FullTimestamp: true, 
		ForceColors: true, })
	return logger
}

func LoggerToFile(ph string) gin.HandlerFunc {
	logger := Logger(ph)

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		logger.Info(" status=%3d  latencyTime=%9s  Address=%15s  method=%s  URI=%s ",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}

/*
###############################################
*/
func Logger2() *logrus.Logger {
	//now := time.Now()
	
	config.Init()
	//conf.App.ServerLog()
	logFilePath := config.GetServerLogPath()
	logFileName := "ph"
	//logFileName := config.GetUrl1LogName()
	//if dir, err := os.Getwd(); err == nil {
	//	logFilePath = dir + "/logs/"
	//}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
	}
	//logFileName := now.Format("2006-01-02") + ".log"

	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}

	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	logger2 := logrus.New()
	logger2.Out = src
	logger2.SetLevel(logrus.DebugLevel)
	logger2.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-1-2 15:04:05.9999",FullTimestamp: true, 
		ForceColors: true, })
	return logger2
}


func URILoggerToFile() gin.HandlerFunc {
	logger2 := Logger2()
	return func(c *gin.Context) {
		startTime := time.Now()
		//time.Now()
		c.Next()
		//fmt.Println(startTime)
		endTime := time.Now()
		//fmt.Println(endTime)
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		logger2.Infof(" %3d  %13s  %15s  %s  %s ",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}