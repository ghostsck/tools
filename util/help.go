package util

import (
	"bytes"
	"fmt"
	guuid "github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"time"
)

/**
CreateUuid
*/
func CreateUuid() string {
	id := guuid.New()
	return id.String()
}

func CreatePasswd(password string) (string, bool) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err == nil {
		encodePW := string(hash)
		return encodePW, true
	} else {
		return "", false
	}
}

func CheckPasswd(checkPasswd string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(checkPasswd), []byte(password))
	if err == nil {
		return true
	} else {
		return false
	}
}

func SetPasswd(passwd, random string) string {
	var buffer bytes.Buffer
	buffer.WriteString(passwd)
	buffer.WriteString(random)
	passwdRandom := buffer.String()
	return passwdRandom
}


func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}


func CreateRandNumber(number int) string {
	numberString := strconv.Itoa(number)
	return fmt.Sprintf("%0"+numberString+"v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

//时间戳
func GetTime() int64 {
	return (time.Now().Unix())
}

//当前时间
func GetDate() string {
	// 获取指定时间戳的年月日，小时分钟秒
	t := time.Unix(GetTime(), 0)
	return fmt.Sprintf("%d-%d-%d %d:%d:%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func ToDate() string {
	// 获取指定时间戳的年月日，小时分钟秒
	t := time.Unix(GetTime(), 0)
	return fmt.Sprintf("%d%d%d", t.Year(), t.Month(), t.Day())
}

func ToInt(num string) int {
	number, _ := strconv.Atoi(num)
	return number
}

func GetSumBytes(bytes int64) int64 {
	var (
		kb  int64 = 1024
		mb  int64 = 0
		gb  int64 = 0
		tb  int64 = 0
		sum int64 = 0
	)

	mb = kb * 1024
	gb = mb * 1024
	tb = gb * 1024

	if bytes < kb {
		sum = bytes
	} else if bytes < mb {
		sum = bytes / kb
	} else if bytes < gb {
		sum = bytes / mb
	} else if bytes < tb {
		sum = bytes / gb
	}

	return sum
}

func GetBytes(number int64, inType string) int64 {
	var sum int64 = 0

	if inType == "kb" {
		sum = number * 1024
	} else if inType == "mb" {
		sum = (1024 * 1024) * number
	} else if inType == "gb" {
		sum = (1024 * 1024 * 1024) * number
	} else if inType == "tb" {
		sum = (1024 * 1024 * 1024 * 1024) * number
	}
	return sum
}
