package utils

import (
	"strings"
)

/*
根据客户端语言环境，取得相关语言字符串
key
lang：取自客户端的Accept-language
*/
func Lang(key, lang string) string {
	return Tr(key, Local(lang))
}

/*
根据客户端语言环境，确定客户端语言
lang：取自客户端的Accept-language
*/
func Local(lang string) string {
	switch {
	case strings.Contains(lang, "zh"):
		return "zh"
	case strings.Contains(lang, "en"):
		return "en"
	default:
		return "zh"
	}
}
