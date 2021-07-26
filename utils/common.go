package utils

import (
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"log"
	"os"
	"regexp"
)

//邮箱验证
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

//字段加密
func ScryptStr(str string) string {
	const keyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 34, 222, 64, 23, 12, 56, 76}
	HashStr, err := scrypt.Key([]byte(str), salt, 16384, 8, 1, keyLen)
	if err != nil {
		log.Fatal(err)
	}
	fstr := base64.StdEncoding.EncodeToString(HashStr)
	return fstr
}

//判断文件是否存在

func IsFileExist(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true

}
