package utils

import (
//"fmt"
)

//列表是否包含给定项
func ListContains(list []interface{}, key interface{}) (finded bool) {

	for _, v := range list {
		if v == key {
			finded = true
			break
		}
	}
	return
}

//字符串数组中是否包含给定项
func StringsContains(list []string, key string) (finded bool) {
	for _, v := range list {
		if v == key {
			finded = true
			break
		}
	}
	return
}
