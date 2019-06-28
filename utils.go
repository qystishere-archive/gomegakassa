package gomegakassa

import (
	cmd5 "crypto/md5"
	"fmt"
)

func md5(s string) string {
	return fmt.Sprintf("%x", cmd5.Sum([]byte(s)))
}
