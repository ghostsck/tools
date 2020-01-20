package main

import (
	"fmt"
	"tools/authcode"
	"tools/util"
)

func main() {

	key := util.EncodeMD5("123456")
	en, _ := authcode.Encrypt("qwertyuiopasdfghjkzxcvbn!@#$%^&*+__^##$%^&*(!@#$%ERTSDFGHVBSERTYHFGHFGTYHGYHG", key, 0)
	fmt.Println(en)

	de, _ := authcode.Decrypt(en, key)
	fmt.Println(de)
}
