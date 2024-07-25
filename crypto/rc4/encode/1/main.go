package main

import (
	"crypto/rc4"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

func rc4encode(shellcode []byte, key []byte) []byte {
	cipher, err := rc4.NewCipher(key)
	if err != nil {

	}
	encryptedBytes := make([]byte, len(shellcode))
	cipher.XORKeyStream(encryptedBytes, shellcode)
	return encryptedBytes
}

func getCurrentYearAsBytes() []byte {
	currentTime := time.Now()

	// 获取当前年份
	currentYear := currentTime.Year()

	// 将年份转换为字符串
	yearStr := strconv.Itoa(currentYear)

	// 将字符串转换为字节数组
	key := []byte(yearStr)

	return key
}

func main() {
	key := getCurrentYearAsBytes()
	eccode := "shellcode"
	data, _ := hex.DecodeString(eccode)
	s := rc4encode(data, key)
	fmt.Printf("加密密钥: %v \n", string(key))
	fmt.Printf("加密数据: %v \n", string(data))
	fmt.Printf("加密结果: %v \n", string(s))
}
