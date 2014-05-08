package utils

import (
	"strconv"
)

//字符串转长整型
func Str2int64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

//字符串转整形
func Str2int(s string) (int, error) {
	return strconv.Atoi(s)
}

//整形转字符串
func Int2str(i int) string {
	return strconv.Itoa(i)
}
