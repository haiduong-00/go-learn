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
- Snippet để test trong VSCode là tf

- t.Error(): Báo ra là test FAIL, chạy tiếp đoạn dưới (continue)
- go command đã học:
    + go test: để chạy tất cả đoạn test
    + go test -v: để theo dõi quá trình các test hoạt động (cho ra nhiều thông tin hơn)

- Test không chỉ là 1 test riêng biệt mà còn có thể là một bảng test: và testqua từng test con (Table test)
- Trong VSCode có sẵn snippet cho điều này, và là tdt

Viết Example cho Documentation
- Quy tắc viết ví dụ cho Documentation
    + Example nó cũng như là 1 test bình thường, nhưng sẽ được sử dụng để mô tả về chức năng của function, hay type đấy
    + Example phải có tên là ExampleXxx với Xxx là tên function/package/type/method cần mô tả ví dụ, và X viết hoa chữ cái đầu
- Snippet cho example là ef

Golint
- Chỉ ra lỗi hay cách sử dụng code Go chưa hợp lý
- Cách dùng: go command:
    + golint . : lint qua các file; ./... lint qua các file của thư mục đó
- go fmt, go vet, golint:
    + go fmt: format code lại cho dễ nhìn và đưa lỗi nếu có trường hợp không ổn
    + go vet: đưa ra những cách dùng sai mà nghiêm trọng đến code Go
    + golint: đưa ra những cách dùng chưa idiomatic (chuẩn nhất)