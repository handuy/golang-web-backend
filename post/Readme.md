# Chaỵ lệnh

```go
go run main.go
```

## Web server khới động tại địa chỉ: [http://localhost:8080/](http://localhost:8080/ "http://localhost:8080/")

## Truy cập vào địa chỉ: [http://localhost:8080/](http://localhost:8080/ "http://localhost:8080/"). Web server trả về form tính chỉ số BMI

![Trang chủ](form.png?raw=true "Trang chủ")

## Khi click Submit, client gửi một POST request lên địa chỉ [http://localhost:8080/upload-form](http://localhost:8080/upload-form "http://localhost:8080/upload-form"). Server khi nhận được request sẽ đọc dữ liệu từ thuộc tính name của các thẻ <input> trong form gửi lên và trả về kết quả

![Kết quả BMI](result.png?raw=true "Kết quả BMI")