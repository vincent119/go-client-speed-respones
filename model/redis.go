package model

import (
	"context"
	"fmt"

	"github.com/vincent119/go-client-speed-respones/config"

	//"github.com/vincent119/go-client-speed-respones/loggin"
	//"github.com/go-redis/redis/v8"
	"github.com/gomodule/redigo/redis"
)

var ctx = context.Background()
var rc redis.Conn

//var rdb *redis.Client

func RedisConnection() redis.Conn {

	/*return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			IPPort := fmt.Sprintf("%s:%s", config.RedisHost(), config.RedisPort())
			rc, err := redis.Dial("tcp", IPPort, redis.DialPassword(config.RedisAuth()))
			//rc, err = redis.Do("AUTH", fmt.Sprintf("%s", config.RedisAuth()))
			if err != nil {
				panic(err.Error())
			}
			return rc, err
		},
	}*/
  IPPort := fmt.Sprintf("%s:%s", config.RedisHost(), config.RedisPort())
  rc, err := redis.Dial("tcp", IPPort, redis.DialPassword(config.RedisAuth()))
	if err != nil {
		panic(err)
	}
	return rc
}


/* func RedisInit() (err error) {
	//config.Init()
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.RedisHost(), config.RedisPort()),
		Password: fmt.Sprintf("%s", config.RedisAuth()),
		DB:       0,
		PoolSize: 5,
	})
	result := rdb.Ping(context.Background())
	fmt.Println("redis ping:", result.Val())
	if result.Val() != "PONG" {
		// connecttion fail
		fmt.Println("redis not connect.........")
		panic(err)
	}
	fmt.Println("connect Redis succeed.........")
	//	_, err = rdb.Ping(context.Background()).Result()
	//return err
	return
}
func RdbGet() {
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}
*/
/*
func RedisGet() {
  fmt.Println("11111111111111111")
  var rdb *redis.Client
  err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
  val, err := rdb.Get(ctx, "key").Result()
  fmt.Println(val)
} */
