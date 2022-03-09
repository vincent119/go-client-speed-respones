package token
import (
	//"github.com/vincent119/go-client-speed-respones/config"
	"fmt"
	"crypto/sha256"
)

func GenSha256(x string) string {
	hash := sha256.Sum256([]byte(x))
	return fmt.Sprintf("%x", hash)
} 