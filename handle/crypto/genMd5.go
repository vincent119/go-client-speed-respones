package crypto

import (
	"crypto/md5"
	//"github.com/vincent119/go-client-speed-respones/config"
	"fmt"
)

func GenMd5 (x string) string {
//	config.GetServerUkey()
	//config.GetServerSalt()
	//data := []byte(x)
	has := md5.Sum([]byte(x))
  return fmt.Sprintf("%x", has)
}