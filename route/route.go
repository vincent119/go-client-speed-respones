package initialize

import (

	"github.com/gin-gonic/gin"
	"net/url"
	"strconv"
)


func InitRouters(r *gin.Enging)  {
	r.POST("sn",ApiVer)
	GroupV1 := r.Group("/v1"){
		GroupV1.Any("/clinetrsp.")

	}
	retune Router
 }

 func ApiVer(c *gin.Context) {




 }