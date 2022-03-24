package token

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vincent119/go-client-speed-respones/config"
	"github.com/vincent119/go-client-speed-respones/handle/cache"
	co "github.com/vincent119/go-client-speed-respones/handle/crypto"
	hp "github.com/vincent119/go-client-speed-respones/handle/http"
)

func CheckHttpToken(c *gin.Context) {
	//TokenValues := hp.HeaderUtoken(c)
	if config.GetServerUkey() != hp.HeaderUtoken(c) {
		http401Return(c)
	}
}

func CheckHttpXkey(c *gin.Context) {
	xkey := hp.HeaderXkey(c)
	if xkey == "" {
		fmt.Println("xkey not found.....")
		http401Return(c)
}}

func CheckRdbXkey(c *gin.Context) {
	count := 0
  xkey := hp.HeaderXkey(c)
	//fmt.Printf("xkey: %s\n",xkey)
	rs := cache.RedisGet(xkey)
	//fmt.Printf("rs: %s\n",rs)
	if rs == "" {
		fmt.Println("redis key value not found......")
		count= count+1
	}
	md5ValueS := xkey + ":" + config.GetServerSalt()
	fmt.Println(co.GenSha256(md5ValueS))
	if rs != co.GenSha256(md5ValueS) {
		fmt.Println("value not match....")
		count= count+1
	}
	fmt.Println(count)
	if count >= 1 {
		http401Return(c)
	} else{
    c.Next()
	}
}

func http401Return(c *gin.Context) {
	c.JSON(401, gin.H{
		"Status": "401",
	})
	c.Abort()
}
