# Chaỵ lệnh

```go
go run main.go
```

## Lưu ý: Ví dụ này không sử dụng Go-gin web framework, chỉ sử dụng package net/http có sẵn trong Go

## Web server khới động tại địa chỉ: [http://localhost:8080/](http://localhost:8080/ "http://localhost:8080/")

## Truy cập vào địa chỉ: [http://localhost:8080/](http://localhost:8080/ "http://localhost:8080/"). Browser hiển thị

![Màn hình browser](browser.png?raw=true "Màn hình browser")

## Trên terminal hiển thị các dòng chữ theo trình tự như sau:

![Màn hình terminal](terminal.png?raw=true "Màn hình terminal")

------------

## Giải thích code

### Trước tiên, định nghĩa một kiểu dữ liệu là Middleware: Là một hàm có tham số đầu vào là `http.HandlerFunc` và trả về một `http.HandlerFunc` khác

### Hàm LogPath không nhận tham số đầu vào và trả về dữ liệu kiểu Middleware

```go
func LogPath() Middleware {
    log.Println("Chạy hàm LogPath")
    // Trả về Middleware: Một hàm nhận http.HandlerFunc và trả về http.HandlerFunc
    return func(f http.HandlerFunc) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
            log.Println("Tên đường dẫn:", r.URL.Path)
            f(w, r)
        }
    }
}
```

### Hàm LogMethod cũng tương tự hàm LogPath

```go
func LogMethod() Middleware {
    log.Println("Chạy hàm LogMethod")
    return func(f http.HandlerFunc) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
            log.Println("Tên method:", r.Method)
            f(w, r)
        }
    }
}
```

### Hai hàm LogPath và LogMethod đều có thêm 1 đoạn log ở đầu hàm để tiện kiểm tra thứ tự chạy hàm

### Hàm Chain là một variadic function gồm các tham số: một biến kiểu `http.HandlerFunc` và các biến kiểu `Middleware`

```go
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
    log.Println("Chạy hàm Chain")
    for _, m := range middlewares {
        f = m(f)
    }
    return f
}
```

### Nhiệm vụ của hàm Chain: Chạy một vòng lặp qua tất cả các Middleware. Ở mỗi lần chạy, Middleware nhận vào `http.HandlerFunc` và trả về một ``http.HandlerFunc` mới

### Hàm Hello là một `http.HandlerFunc` thông thường

```go
func Hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "<h1>Hello World</h1>")
}
```

### Khi khởi động web server, server sẽ lần lượt chạy các hàm theo thứ tự

1. Chạy hàm LogPath
2. Chạy hàm LogMethod
3. Chạy hàm Chain

### Sau khi chạy 2 hàm LogPath và LogMethod, lúc này đã có 2 Middleware sẵn sàng nhận `http.HandlerFunc` để trả về một `http.HandlerFunc` mới. Đến khi chạy hàm Chain, 2 Middleware vừa mới tạo sẽ lần lượt nhận hàm Hello theo thứ tự

1. Hàm LogPath nhận vào hàm Hello và trả về một `http.HandlerFunc` mới, tạm gọi là Hello-1. Hàm Hello-1 mới này sẽ log ra terminal tên URL path, trả về dòng chữ "Chạy hàm LogPath" trên browser, sau đó mới chạy hàm Hello

2. Hàm LogMethod nhận vào ***hàm Hello-1 vừa mới được tạo ở bước 1*** và trả về một `http.HandlerFunc` mới, tạm gọi là Hello-2. Hàm Hello-2 sẽ log ra terminal tên method, trả về dòng chữ "Chạy hàm LogMethod" trên browser, sau đó mới chạy hàm Hello-1

### Như vậy, khi có request gửi đến /, server sẽ xử lý theo thứ tự

1. Trả về dòng chữ "Chạy hàm LogMethod" cho browser

2. Trả về dòng chữ "Chạy hàm LogPath" cho browser

3. Chạy hàm Hello: Trả về dòng chữ "Hello World" cho browser