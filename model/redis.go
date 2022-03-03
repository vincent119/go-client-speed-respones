package model

import (
  "context"
  "fmt"
  "github.com/vincent119/go-client-speed-respones/config"
  //"github.com/vincent119/go-client-speed-respones/loggin"
  "github.com/go-redis/redis/v8"
)


func RedisInit() *redis.Client {
  //config.Init()
  rdb := redis.NewClient(&redis.Options{
  Addr: fmt.Sprintf("%s:%s", config.RedisHost(), config.RedisPort()),
  Password: fmt.Sprintf("%s",config.RedisAuth()),
  DB: 0,
    })
  result := rdb.Ping(context.Background())
  fmt.Println("redis ping:", result.Val())
  if result.Val()!="PONG"{
	// connecttion fail
	fmt.Println("redis ping:", result.Val())
	return nil
  }
  return rdb
}


