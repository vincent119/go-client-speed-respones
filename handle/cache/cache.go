package cache

import (
	"github.com/vincent119/go-client-speed-respones/config"
	rds "github.com/vincent119/go-client-speed-respones/handle/rdsub"
)

//var rc = model.RedisConnection()

func RedisSet(md5Value string, sha256Value string) {
	rds.Set(md5Value, sha256Value, config.RedisTtl())
}

func RedisGet(key string) string {
	rep  := rds.Get(key)
	return rep
}
