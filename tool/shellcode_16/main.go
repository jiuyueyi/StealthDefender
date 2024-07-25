package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}
	filename := os.Args[1]

	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 读取文件内容
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 正则表达式匹配花括号中的内容
	re := regexp.MustCompile(`\{([^}]*)\}`)
	matches := re.FindAllStringSubmatch(string(content), -1)

	// 提取并处理匹配到的内容
	for _, match := range matches {
		if len(match) > 1 {
			trimmedStr := strings.TrimSpace(match[1])
			trimmedStr = strings.ReplaceAll(trimmedStr, ",", "")
			trimmedStr = strings.ReplaceAll(trimmedStr, " ", "")
			trimmedStr = strings.ReplaceAll(trimmedStr, "0x", "")
			fmt.Println(trimmedStr)
		}
	}
}
