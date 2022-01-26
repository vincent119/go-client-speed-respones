package main

import (
	"fmt"
	"os"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	//"go-client-speed-respones/model"
)


func init(){
  confPath ,_ := os.Getwd()  
	viper.AddConfigPath("./config/")
	viper.AddConfigPath(confPath)
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
  	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("no such config file")
	} else {
			// Config file was found but another error was produced
			log.Println("read config error")
	}
	log.Fatal(err) 
  }
}

func main(){
	fmt.Println("application port = " + viper.GetString("app.port"))
  fmt.Println(viper.GetString("app.logfile"))


  r := gin.Default()
	r.SetTrustedProxies([]string{"172.16.99.200"})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
				"message": "pong",
		})
	})
	r.Run()
}