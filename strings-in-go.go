package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	//byte: 8bits: 2^8 states
	//ASCII 7bits: 2^7 states -> 2^7 characters
	//String: một chuỗi byte (byte=8bits) //Ngôn ngữ lập trình, ngoài đời: String là một chuỗi ký tự
	//Hello世界: H(byte 0) e(byte 1) l(byte 2) l(byte 3) o(byte 4) 世(byte 5, 6, 7) 界(byte 8, 9, 10)
	//hexadecimal: binary(2) hexadecimal=hex(16)  decimal(10):UTF-8
	//Strings: Go chủ yếu là những ký tự của UTF-8, Go dịch được mỗi UTF-8
	//UTF-8: 1 ký tự có thể chiếm đến 4 byte: 2^16 states, 2^16 codepoint
	//Strings are immutable
	//s[i]: byte ở vị trí thứ i,+ trả về giá trị binary
	//rune(trong UTF-8 codepoint) == int32: vì UTF-8 có kí tự chạy từ hex(16) U+0000 -> U+FFFF
	s := "\xc2\xbdHello, 世界" //thêm byte C2
	fmt.Println(s)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v %+U\n", i, s[i])
	}
	fmt.Println()
	for i, v := range s {
		fmt.Printf("%v %+U\n", i, v) //i: index of starting byte
	}
}
