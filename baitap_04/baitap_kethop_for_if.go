//Cho một số a bất kì, a thuộc tập N*, thực hiện những điều sau trong vòng lặp
//Nếu a là số chẵn, chia 2 cho chính nó và tiếp tục vòng lặp
//Nếu a là số lẻ, nhân a với 3 và cộng thêm 1 và tiếp tục vòng lặp
//Nếu a == 1 thì thoát khỏi vòng lặp
//Cuối cùng in ra số lần lặp

//Vấn đề nào cũng có đề bài cả (Con cá nguyên chất)
//Đề bài là vật liệu thô
//Thuật toán là đã sơ chế (Con cá đã làm vẩy, ướp muối)
//Những thứ cần thiết cho đoạn code (Gia vị)
//Nó cần được chế biến thành đồ dùng được, đồ ăn được

//for, a != 1, if, %:mod: chia lấy dư, a(biến), i :biến, nó có thể được truy cập ở ngoài vòng for
package main

import (
	"fmt"
)

func main() {
	i := 0
	a := 22
	for a != 1 {
		if a%2 == 0 {
			a = a / 2
			i++
		} else {
			a = a*3 + 1
			i++
		}
	}
	fmt.Println(i)
}
