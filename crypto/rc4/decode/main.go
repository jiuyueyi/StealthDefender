package main

import (
	"crypto/rc4"
	"fmt"
	"log"
	"strconv"
	"time"
)

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

func rc4decode(shellcode []byte, key []byte) []byte {
	cipher, err := rc4.NewCipher(key)
	if err != nil {
		log.Println(err)
	}
	decryptedBytes := make([]byte, len(shellcode))
	cipher.XORKeyStream(decryptedBytes, shellcode)
	return decryptedBytes
}

func main() {
	key := getCurrentYearAsBytes()
	s := []byte{123, 23} //shellcode
	d := rc4decode(s, key)
	fmt.Printf("解密结果: %v \n", string(d))
}
