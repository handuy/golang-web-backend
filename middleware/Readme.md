# Chaỵ lệnh

```go
go run main.go
```

## Lưu ý: Ví dụ này không sử dụng Go-gin web framework, chỉ sử dụng package net/http có sẵn trong Go.

## Web server khới động tại địa chỉ: [http://localhost:8080/](http://localhost:8080/ "http://localhost:8080/"). Màn hình terminal hiển thị 2 dòng chữ: Chạy hàm logging

![Màn hình terminal](gorun.png?raw=true "Màn hình terminal")

## Điều này chứng tỏ khi khởi động, web server sẽ chạy hàm logging trước. Hàm logging được gọi ở 2 đường dẫn: /foo và /bar, do đó màn hình terminal mới hiển thị 2 dòng chữ "Chạy hàm logging"

![Logging function](logging.png?raw=true "Logging function")

## Hàm logging sau khi chạy xong sẽ trả về một function dạng http.HandlerFunc. Function này có nhiệm vụ xử lý request gọi đến /foo và /bar

![Route request handler](middleware_func.png?raw=true "Route request handler")

## Truy cập vào địa chỉ: [http://localhost:8080/foo](http://localhost:8080/foo "http://localhost:8080/foo"). Web server trả về chữ "foo" in trên browser, đồng thời trên terminal cũng hiển thị 2 dòng chữ: "Chạy hàm middleware" và tên đường dẫn "/foo"

![Foo browser](foo_browser.png?raw=true "Foo browser")

![Foo terminal](foo_terminal.png?raw=true "Foo terminal")

## Gọi vào địa chỉ [http://localhost:8080/bar](http://localhost:8080/bar "http://localhost:8080/bar") cũng cho kết quả tương tự

![Bar browser](bar_browser.png?raw=true "Bar browser")

![Bar terminal](bar_terminal.png?raw=true "Bar terminal")

## Điều này cho thấy tác dụng của middleware logging: Khi có request gửi đến /foo và /bar, trước khi chạy vào hàm xử lý chính foo và bar, request sẽ được xử lý bởi một hàm middleware: Hàm này sẽ in ra dòng chữ "Chạy hàm middleware" và tên đường dẫn trên terminal

## Các bạn có thể không sử dụng middleware logging bằng cách gọi trực tiếp đến 2 hàm xử lý chính foo và bar, sau đó so sánh kết quả

![No middleware](no_middlware.png?raw=true "No middleware")