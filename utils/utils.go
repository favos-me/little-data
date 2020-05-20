package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(origin string) (str string) {
	hasher := md5.New()
	hasher.Write([]byte(origin))
	return hex.EncodeToString(hasher.Sum(nil))
}

//数组去重
func RemoveDuplicate(slis []string) {
	found := make(map[string]bool)
	j := 0
	for i, val := range slis {
		if _, ok := found[val]; !ok {
			found[val] = true
			(slis)[j] = (slis)[i]
			j++
		}
	}
	slis = (slis)[:j]
	return
}
