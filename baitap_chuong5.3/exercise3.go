package main
import("fmt")
func main(){
	x:=[]int{42,43,44,45,46,47,48,49,50,51}
	fmt.Println(x[:5])
	slicingSlice := func (l,r int) []int {
		return x[l:r]
		// return sẽ yêu cầu 1 expression (operands và operators (số hạng) và (dấu))
		// statement: identifier := expression
	}
	fmt.Println(slicingSlice(0,5))
	// fmt.Println(x[2:7])
	// fmt.Println(x[1:6])
}

// Viết func để cắt 1 slice như trên [l:r]



// cắt x[4:5] -> {46,47}

// b = a
// a[1] = 4
// b[1] = 2 -> a[1] = 2
