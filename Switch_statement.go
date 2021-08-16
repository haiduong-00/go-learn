//Trong Go, hay ở các ngôn ngữ khác đều dạng điều kiện if
//Và còn loại conditional thứ 2: Switch (Conditional: nếu theo đúng điều kiện thì chạy đoạn code, không thì thôi)

//Cách tạo nên câu switch:
/*switch expression {
case condition:

}*/

//Keyword (những từ mà ngôn ngữ đã chiếm) của switch tiếp theo là fallthrough
//Fallthrough: nếu cái điều kiện chứa nó đúng: thực hiện case tiếp theo (bỏ qua cả diều kiện của case)

//Keyword: default: mặc định: nếu các cái case không có cái nào thỏa mãn, thì default case sẽ chạy

package main

import "fmt"

func main() {
	//i := 24
	/*switch {
	default:
		fmt.Println("Không có cái nào thỏa mãn")
	case (i < 9), (i > 23), (i == 25):
		fmt.Println("i nhỏ hơn 9 hoặc lớn hơn 23")
		fallthrogh
	case i < 11:
		fmt.Println("i lớn hơn 11")
		fallthrough
	case i < 20:
		fmt.Println("i lớn hơn 20")
	case i == 10:
		fmt.Println("i bằng 10")
	}

	if false == (i < 9) { // true == (false == (i < 9))
		fmt.Println("Hi")
	}
	if false == ((i < 9) || (i < 23) || (i == 25)) {
		fmt.Println("i nhỏ hơn 9 hoặc lớn hơn 23")
	} */
	var s string = "24"
	switch i := 0; s {
	case "Hello", "hi", "My Name":
		fmt.Println("run this", i)
		fallthrough
	case "Bond", "James", "Monneypenny":
		fmt.Println("run that", i)
	}

	/*Đang trong chat
	switch {
	case co tin nhan: //Không
		rep()
		fallthrough //đưa mình dến hành động phía dưới luôn
	case muốn chờ tin nhắn không:
		chotinnhan(2s)
	} */

	//Logic lập trình for if //nếu là trường hợp: thì dùng if, nếu có biến chạy, số chạy hay một thứ gì đó lặp đi lặp lại: for, switch, logic toán học: false true
	//if lồng trong for: trong khi chạy for, mà đạt đến điều kiện if thì thực hiện cái gì đó
	//switch case == if nhưng ở dạng viết khác, cùng một mẹ Conditional

	//Composite data type: Object/Struct, Array, Slices, Map: 2 tuần
}
