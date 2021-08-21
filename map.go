package main

import "fmt"

func main() {
	mSaveAge := map[string]string{
		"James":      `32`,
		"Moneypenny": `28`,
	}
	// map[<type của key>]<type của value>
	// Vì map không có chỉ số (index), nên không thể sử dụng slicing
	// cũng không thể append (vì append chỉ dành cho slice)
	// Nhưng mà có một thứ: for range
	// có thể sử dụng ở map, slice, array, channel
	fmt.Println(mSaveAge)
	for k := range mSaveAge { //range thường trả về 2 giá trị: index/key - value
		fmt.Println(k,mSaveAge[k])
	}
	// Muốn thêm giá trị, thì thêm 1 cặp key-value
	mSaveAge["Todd"] = `36`
	fmt.Println("todd", mSaveAge["todd"]) // Default value, zero value
	// Muốn xóa 1 key nào đó: 
	delete(mSaveAge, "James")
	fmt.Println(mSaveAge)

}
