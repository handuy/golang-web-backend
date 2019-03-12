# Chaỵ lệnh

```go
go run main.go
```

## Web server khới động tại địa chỉ: [http://localhost:8080/](http://localhost:8080/ "http://localhost:8080/")

## Truy cập vào địa chỉ: [http://localhost:8080/](http://localhost:8080/ "http://localhost:8080/"). Web server trả về form upload audio. Form này cho phép upload nhiều file cùng lúc

![Form upload audio](upload-mp3.png?raw=true "Form upload audio")

## Sau khi chọn file audio xong và click Submit, client gửi một POST request lên địa chỉ [http://localhost:8080/upload-audio](http://localhost:8080/upload-image "http://localhost:8080/upload-image").

## Server khi nhận được POST request sẽ tiến hành đọc dữ liệu các file gửi lên và lưu vào thư mục public. Nếu lưu thành công sẽ trả về file JSON là một mảng tên các file được upload lên. Nếu thất bại sẽ redirect sang đường dẫn /fail

![Upload audio thành công](upload-success.png?raw=true "Upload audio thành công")