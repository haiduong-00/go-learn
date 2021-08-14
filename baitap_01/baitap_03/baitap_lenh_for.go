//In ra số từ 65 đén 90
//Chuyển những số đó sang dạng ký tự (gợi ý: sử dụng %c)

package main

import (
	"fmt"
)

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%c", i)
	}
}
