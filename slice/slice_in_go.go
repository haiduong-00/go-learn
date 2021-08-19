package main
import ("fmt")
func main() {
	// x:= []int{} // { }: mà đứng sau type of slice, composite literal (Giá trị sẽ gán sẵn, là hằng số): Nếu tạo slice bằng := (short declaration operator), hay gán giá trị cho cả slice, composite literal: bắt buộc 
	// var xint []int  // Phần []int bắt buộc
	 //Tạo slice với 100 giá trị ban đầu = 0 (Zero value)
	// Sử dụng make
	xint := make([]int,10000000)
	// xint = append(xint, 101) //Sử dụng append: thêm giá trị
	// xint = xint[:200]  //Slice (cắt) một slice  : Là mày (máy tính) phải tăng cái cap lên gấp đôi
	// Trong array, length == capacity
	// Trong slice, length != capacity
	xint[101] = 25
	fmt.Println(len(xint),cap(xint))
}