//file:base64.go
//auth:lip
//date:2018-09-27
//desc:base64加解密
package utils

import (
	"encoding/base64"
	"fmt"
)

//base64加密
func Baese64Encode(str string) string {
	strbytes := []byte(str)
	encoded := base64.StdEncoding.EncodeToString(strbytes)
	return encoded
}

//base64解密
func Base64Decode(encoded string) string {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("解密失败")
	}
	decodestr := string(decoded)
	return decodestr
}
