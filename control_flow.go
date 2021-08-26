package main

import (
	"fmt"
)

var y int

func main() {
	//Ứng dụng Go, C++ sẽ đọc (ở trong codeblock "{ }") từ trên xuống dưới, từ trái qua phải
	//Còn ở ngoài codeblock main nó đọc thế nào cũng được, và nó được đọc trước tiên khi vào chương trình
	//Ở ngoài là khâu chuẩn bị sẵn sàng cho func main() hay codeblock khác
	//sequence/flow of computer program
	//JavaScript: env: environment
	//Trong JS: mỗi cái func là một env, còn trong GO: codeblock là một env
	//Trong một chương trình: sẽ có ít nhất 1 trong 3 phần
	
	//Sequential: dòng chảy bình thường

	y = 9
	fmt.Println(y)
	i := 0

	//Loop/Iterative: vòng lặp

	for i = 0; i < 10; i++ {
		fmt.Println(i)
	}

	//Conditional: câu điều kiện
	fmt.Println("i trước", i)
	if i > y {
		i -= 2
		fmt.Println("i sau:", i)
	}
}

