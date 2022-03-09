package token

import (
	"crypto/md5"
	"github.com/vincent119/go-client-speed-respones/config"
	"fmt"
)

func GenMd5 (x string) string {
	config.GetServerUkey()
	config.GetServerSalt()
	data := []byte(config.GetServerUkey()+config.GetServerSalt())
	has := md5.Sum(data)
  return fmt.Sprintf("%x", has)
}