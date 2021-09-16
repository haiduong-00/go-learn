package main

// Context: 1 model, concept về concurrency: giúp cho việc giới hạn giá trị, kết quả, và resources (tài nguyên máy tính) sao cho hệ thống concurrency an toàn và hiệu quả nhất
// Khi mà sử dụng context: thì đã gói cái hành động, quá trình nào đấy riêng biệt ra
// Ví dụ: khi đóng gói hàng hóa: Quá trình đóng gói cái máy tính sẽ ở trong 1 context, và nó thực hiện những routine liên quan nó, những yêu cầu, kết quả đều trong một context, Context gói máy tính này TÁCH BIỆT với context gói cái đồng hồ, không liên kết gì với nhau
// Trong context có 3 hành động: WithCancel: Cancel context bất cứ lúc nào nếu muốn điều kiện
// WithDeadline (Nó nêu thời gian cụ thể vd: 9h tối, 7h sáng), WithTimeout (Nó đưa ra khoảng thời gian vd: 15 phút, trong 1 tiếng): Cancel context khi đạt đến thời gian nào đó với trả về điều kiện

import (
	"context"
	"fmt"
)

func main() {
	con := context.Background()
	con,cancel := context.WithCancel(con)
	cancel()
	fmt.Printf("%T %v %v",con,con,con.Err())
}
