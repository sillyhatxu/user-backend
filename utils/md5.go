package utils

import (
	"crypto/md5"
	"fmt"
)

func MD5(src string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(src)))
}
