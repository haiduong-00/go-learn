package main

import (
	"fmt"
)

func main() {
	// x:= []int{} // { }: mà đứng sau type of slice, composite literal (Giá trị sẽ gán sẵn, là hằng số): Nếu tạo slice bằng := (short declaration operator), hay gán giá trị cho cả slice, composite literal: bắt buộc
	// var xint []int  // Phần []int bắt buộc
	//Tạo slice với 100 giá trị ban đầu = 0 (Zero value)
	// Sử dụng make

	// xint = append(xint, 101) //Sử dụng append: thêm giá trị
	// xint = xint[:200]  //Slice (cắt) một slice  : Là mày (máy tính) phải tăng cái cap lên gấp đôi
	// Trong array, length == capacity
	// Trong slice, length != capacity
	xint := []rune{0, 1, 3, 3, 4, 5, 6, 7}
	xint[2] = 2
	xAdding := []rune{20, 21, 22}
	xAfterAppend := append(xint, 8, 9, 10) // append(<cái slice gốc>, các giá trị thêm vào)
	// Trả về 1 slice mới
	xAfterAppend = append(xAfterAppend, xAdding...) // ... (nếu đằng sau 1 slice, array, string, map) extract all values
	// ... nếu ở phía trước: Nhập bao nhiêu giá trị cũng được
	xAfterAppend = append(xAfterAppend, 23, 24)
	// append(xAfterAppend, xAdding[0], xAdding[1], xAdding[2])
	// Nếu đã sử dụng gỡ gói hàng
	fmt.Println(xint)
	fmt.Println(xAfterAppend)

	// for i,v := range xint {     // range là giới hạn cho i:=0 chạy trong len(xint)
	// 	// range trả về 2 giá trị: i, v (xint[i])
	// 	fmt.Println(i,v)  //index, value
	// }
	// // string là slice of byte
	// s := fmt.Sprintf("Hello, 世界%c%c",0x1f1ef,0x1f1f5)  // slice of UTF-8 rune
	// fmt.Println(len(s))   //13 byte
	// // Đếm ký tự
	// count := 0
	// for _,v := range s {  // range không quy về byte, quy về từng codepoint (rune)
	// 	fmt.Printf("%c\n",v)
	// 	count++
	// }
	// fmt.Println(count)
	// for i:=0 ; i<len(xint);i++{
	// 	fmt.Println(i,xint[i])
	// }
	xAfterSlice := xint[:] // <tên slice>[<vị trí bắt đầu>:<vị trí kết thúc>]
	fmt.Println(xAfterSlice)
	fmt.Println(xint[5]) // Khi cắt slice, vị trí cuối cùng không tính vào
	xAfterDel := append(xAfterAppend[:9],xAfterAppend[10:]...)
	fmt.Println(xAfterDel)

	xCreatebyMake := make([]rune,10,100)
	
	xCreatebyMake = xCreatebyMake[:100]
	fmt.Println(cap(xCreatebyMake))
	xCreatebyMake = append(xCreatebyMake, 100)
	fmt.Println(cap(xCreatebyMake))
	xCreatebyMake = xCreatebyMake[:224]
	xCreatebyMake = append(xCreatebyMake, 100)
	fmt.Println(cap(xCreatebyMake))

	// flappy bird
	// con chim {  object
	// 	dài
	// 	rộng
	// 	rơi
	// 	.onclick(): bay lên bao nhiêu
	// }
	// 	Dog {
	// 		kích thước: công khai {public}
	// 		tên: bí mật (private)
	// 		.sủa
	// 		.giữ nhà
	// 		.ăn 
	// 	}

	// // OOP hiện đại (Go, JavaZscript, Python) != OOP truyền thống (C++, Java)
	// đặc điểm của OOP không đổi: 
	//  cấu trúc, đặc điểm
	//  hành động
	//  kế thừa
	//  công khai, bí mật
}
