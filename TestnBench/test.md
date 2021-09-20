Test

- Kiểm thử: kiểm thử CHỨC NĂNG của đoạn code, xem đoạn code có thực hiện đúng chức năng không, tránh lỗi
- Trong Go: có hệ thống giúp đông việc testing dễ dàng hàng: go test
- Thường đi kèm với code go, sẽ là một code go test
- Code test trong Go là code thủ công: phải tự đánh code test

Hướng dẫn viết code test (quy định code test)
- Các file test sẽ bắt đầu bằng tên có chứa _test.go, và nằm trong cùng package với code chính
- Package file test thường cùng package với file gốc (nếu khác package thì phải import package file gốc)
- Function test chức năng phải có tên là TestXxx với Xxx là tên function cần test, và X viết hoa chữ cái đầu
vd: TestPalindrome để test palindrome
vd: TestHuman_Eat để test method Eat() của struct Human

- t.Error(): Báo ra là test FAIL, chạy tiếp đoạn dưới (continue)
- go command đã học:
    + go test: để chạy tất cả đoạn test
    + go test -v: để theo dõi quá trình các test hoạt động (cho ra nhiều thông tin hơn)

- Test không chỉ là 1 test riêng biệt mà còn có thể là một bảng test: và testqua từng test con