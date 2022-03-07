package cache

import (
	"fmt"
  "os"
  "github.com/gomodule/redigo/redis"
	//"github.com/vincent119/go-client-speed-respones/model"
)

//var rc = model.RedisConnection()

func RedisSet(key string, value string, rc redis.Conn) {
	rc.Do("SET", key, value)
}

func RedisGet(key string, rc redis.Conn) string {
	s, err := redis.String(rc.Do("GET", key))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return s
}