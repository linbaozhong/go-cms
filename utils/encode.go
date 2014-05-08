package utils

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/axgle/mahonia"
	"net/url"
	//"strings"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
)

var key = MD5byte("lin.baozhong19680415")
var iv = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

//md5加密
func MD5(s string) string {
	//是否使用高强度密码
	if b, _ := beego.AppConfig.Bool("StrongPassword"); b {
		return MD5Ex(s)
	} else {
		return hex.EncodeToString(MD5byte(s))
	}
}

func MD5byte(s string) []byte {
	h := md5.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}

//加盐强密码
func MD5Ex(s string) string {
	h := md5.New()
	h.Write(key)
	h.Write([]byte(s))
	h.Write(iv)
	//fmt.Println(hex.EncodeToString(h.Sum(nil)), fmt.Sprintf("%x", h.Sum(nil)), MD5(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

//sha1加密
func SHA1(s string) string {
	return hex.EncodeToString(SHA1Byte(s))
}

func SHA1Byte(s string) []byte {
	h := sha1.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}

//Base64编码
func Base64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

//Base64解码
func Base64Decode(src string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(src)
}

//AES编码
func AesEncode(src []byte) ([]byte, error) {
	var s []byte
	c, err := aes.NewCipher(key)
	if err == nil {
		cfb := cipher.NewCFBEncrypter(c, iv)
		s = make([]byte, len(src))
		cfb.XORKeyStream(s, src)
	}
	return s, err
}

//AES解码
func AesDecode(src []byte) ([]byte, error) {
	var s []byte
	c, err := aes.NewCipher(key)
	if err == nil {
		cfb := cipher.NewCFBDecrypter(c, iv)
		s = make([]byte, len(src))
		cfb.XORKeyStream(s, src)
	}
	return s, err
}

//utf-8转gbk
func Utf8ToGBK(str string) string {
	//字符集转换
	enc := mahonia.NewEncoder("gbk")
	return enc.ConvertString(str)
}

//gbk转utf-8
func GBKToUtf8(str string) string {
	//字符集转换
	enc := mahonia.NewDecoder("gbk")
	return enc.ConvertString(str)
}

//url编码
func UrlEncode(s string) string {
	return url.QueryEscape(s)
}
