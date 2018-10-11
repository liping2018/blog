//file:md5.go
//auth:lip
//date:2018-09-27
//desc:MD5加密

package utils

import (
	"crypto/md5"
	"fmt"
)

//md5加密
func MD5Encode(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}
