package main

import (
	"crypto/rc4"
	"encoding/hex"
	"fmt"
	"github.com/eknkc/basex"
	"log"
)

func main() {
	message := "shellcode"

	//定义base85加密
	base85, err := basex.NewEncoding("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!#$%&()*+-;<=>?@^_`{|}~")
	if err != nil {
		log.Fatalf("Failed to create Base85 encoding: %v", err)
	}

	messageArray, err := hex.DecodeString(message)

	//rc加密
	key := []byte("acb")
	cipher, _ := rc4.NewCipher(key)
	rc4Message := make([]byte, len(messageArray))
	cipher.XORKeyStream(rc4Message, messageArray)
	hexCiphertext := hex.EncodeToString(rc4Message)
	fmt.Printf("rc4 Encoded: %s\n", hexCiphertext)
	//fmt.Printf("rc4Message Encoded Message: %s\n", rc4Message)

	//base85加密
	encodedMessage := base85.Encode(rc4Message)
	//encodedMessage为加密好的shellcode，下面的解密操作只是测试
	fmt.Printf("Base85 Encoded Message: %s\n", encodedMessage)

	//base85解密
	decodedBytes, err := base85.Decode(encodedMessage)
	if err != nil {
		log.Fatalf("Failed to decode Base85 message: %v", err)
	}
	decodedBytes1 := hex.EncodeToString(decodedBytes)
	fmt.Printf("Base85 Decoded: %s\n", decodedBytes1)
	//fmt.Printf("decodedBytes Decoded Message: %s\n", decodedBytes)

	//rc4解密
	cipher, _ = rc4.NewCipher(key)
	xordMessage1 := make([]byte, len(decodedBytes))
	cipher.XORKeyStream(xordMessage1, decodedBytes)
	xordMessage2 := hex.EncodeToString(xordMessage1)
	fmt.Printf("rc4 Decoded: %s\n", xordMessage2)

}
