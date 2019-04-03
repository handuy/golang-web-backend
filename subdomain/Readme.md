Cài đặt kataras iris:
```go
go get -u github.com/kataras/iris
```

Cập nhật file hosts tại: /etc/hosts. Thêm các dòng sau:
```
127.0.0.1       techmaster.local
127.0.0.1       golang.techmaster.local
127.0.0.1       java.techmaster.local

```

Chạy code:
```go
go run main.go
```

Ví dụ gồm:
- 1 Domain chính: techmaster.local:8080
- 2 subdomain: golang.techmaster.local:8080 và java.techmaster.local:8080. Mỗi subdomain có 3 đường dẫn: /, /course và /secret, hiển thị nội dung khác nhau

Khi truy cập vào /secret ở một subdomain, ví dụ golang.techmaster.local:8080/secret:
- Nếu user chưa đăng nhập: Hệ thống trả về HTTP status: Forbidden. Truy cập vào java.techmaster.local:8080/secret cũng bị forbidden
- User truy cập vào golang.techmaster.local:8080/login:  Hệ thống sẽ redirect về trang login của domain chính: techmaster.local:8080/login để user đăng nhập. Server trả về cookie lưu thông tin đăng nhập của user. 
- Sau khi đăng nhập, hệ thống sẽ redirect lại về trang golang.techmaster.local:8080/secret. Lúc này server trả về HTTP 200 OK
- User đã đăng nhập golang.techmaster.local:8080 thì cũng sẽ đăng nhập ở java.techmaster.local:8080. Truy cập vào java.techmaster.local:8080/secret cũng trả về HTTP 200 OK