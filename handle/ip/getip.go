package ip

import (
	"github.com/gin-gonic/gin"


)


func getClientIP(ctx *context.Context) string {
  ip := ctx.Request.Header.Get("X-Forwarded-For")
  if strings.Contains(ip, "127.0.0.1") || ip == "" {
      ip = ctx.Request.Header.Get("X-real-ip")
  }

  if ip == "" {
      return "127.0.0.1"
  }

  return ip
}

// caller
ip := getClientIP(c.Ctx) 