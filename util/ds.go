package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	OS_DS_SALT = "6cqshh5dhw73bzxn20oexa9k516chk7s"
	CN_DS_SALT = "xV8v4Qu54lUKrEYFZkJhB8cuOh9Asafs"
)

func GenDS(str string) string {
	now := time.Now().Unix()
	t := strconv.FormatInt(now, 10)
	r := getRandstring(6)
	h := MD5(fmt.Sprintf("salt=%s&t=%s&r-%s", OS_DS_SALT, t, r))
	return fmt.Sprintf("%s,%s,%s", t, r, h)
}

func getRandstring(length int) string {
	if length < 1 {
		return ""
	}
	char := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charArr := strings.Split(char, "")
	charlen := len(charArr)
	ran := rand.New(rand.NewSource(time.Now().Unix()))

	var rchar string = ""
	for i := 1; i <= length; i++ {
		rchar = rchar + charArr[ran.Intn(charlen)]
	}
	return rchar
}

func MD5(str string) string {
	b := []byte(str)
	h := md5.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}
