package handle
//https://github.com/letseeqiji/gorobbs/blob/master/package/gredis/redis.go
import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	//"log"
	"github.com/vincent119/go-client-speed-respones/config"
	//"github.com/vincent119/go-client-speed-respones/loggin"
	//"github.com/go-redis/redis/v8"
	"github.com/gomodule/redigo/redis"
)

var ctx = context.Background()
var RedisConn *redis.Pool

func init() {
	Setup()

}

func Setup() error {
	RedisConn = &redis.Pool{
		MaxIdle:     3,
		MaxActive:   10,
		IdleTimeout: 0,
		Dial: func() (redis.Conn, error) {
			host := fmt.Sprintf("%s:%s", config.RedisHost(), config.RedisPort())
			c, err := redis.Dial("tcp", host, redis.DialDatabase(0))
			if err != nil {
				return nil, err
			}

			if _, err := c.Do("AUTH", config.RedisAuth); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return nil
}

// Set a key/value
func Set(key string, data interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

// Exists check a key
func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

// Get get a key
func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

// Delete delete a kye
func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

// LikeDeletes batch delete
func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}

func Lpush(key string, value interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	_, err := conn.Do("LPUSH", key, value)
	if err != nil {
		return err
	}
	/*
		_, err = conn.Do("EXPIRE", key, time)
		if err != nil {
			return err
		}*/

	return nil
}

func Brpop(key string) (string, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.String(conn.Do("RPOP", key))
	if err != nil {
		return reply, err
	}

	return reply, nil
}
