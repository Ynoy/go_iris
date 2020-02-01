package utility

import (
	"crypto/md5"
	"encoding/hex"
)

// 获取MD5
func MD5(str string) string {
	_md5 := md5.New()
	_md5.Write([]byte(str))
	return hex.EncodeToString(_md5.Sum([]byte(nil)))
}