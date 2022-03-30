package test

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"git.dustess.com/mk-base/util/rsa"
	mrand "math/rand"
	"time"
)

func init() {
	mrand.Seed(time.Now().UnixNano())
}

type ClientInfo struct {
	LoginPhone string `json:"loginPhone"`
	Domain     string `json:"domain"`
}

const (
	aesKey = "1234567890abcdef"
	aesIv  = "1234567890abcdef"
	// 随机字符串
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

func MyEncrypt(phone, domain string) {
	ci := ClientInfo{
		LoginPhone: phone,
		Domain:     domain,
	}
	b, err := json.Marshal(&ci)
	if err != nil {
		return
	}
	//vi := GetIV()
	k := []byte(aesKey)
	iv := []byte(aesIv)
	res, err := rsa.AesCBCEncrypt(k, b, iv)
	if err != nil {
		return
	}
	s := base64.StdEncoding.EncodeToString(res)
	fmt.Printf("这是加密后的数据 res= [%s]\n iv= [%s]\n vi=[%s]\n", s, base64.StdEncoding.EncodeToString(iv), iv)
	dres := AesCBCDecrypt(k, iv, s)
	fmt.Printf("这是解密后的数据 dres= [%s]\n iv= [%s]\n vi=[%s]\n", dres, string(iv), iv)

}

// AesCBCEncrypt ...
func AesCBCEncrypt(key, data, iv []byte) ([]byte, error) {
	aesBlockEncrypter, err := aes.NewCipher(key)
	content := PKCS5Padding(data, aesBlockEncrypter.BlockSize())
	encrypted := make([]byte, len(content))
	if err != nil {
		println(err.Error())
		return nil, err
	}
	aesEncrypter := cipher.NewCBCEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.CryptBlocks(encrypted, content)
	return encrypted, nil
}

// PKCS5Padding 使用PKCS5进行填充
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS7UnPadding 使用PKCS7进行去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// GetIV 获取IV
func GetIV() string {
	b := make([]byte, 16)
	for i := range b {
		index := mrand.Intn(len(letterBytes))
		b[i] = letterBytes[index]
	}
	return string(b)
}

func AesCBCDecrypt(key, iv []byte, data string) string {
	crytedByte, _ := base64.StdEncoding.DecodeString(data)

	// 分组秘钥
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(fmt.Sprintf("key 长度必须 16/24/32长度: %s", err.Error()))
	}
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, iv)
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return string(orig)
}
