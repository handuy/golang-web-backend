# Chaỵ lệnh

```go
go run main.go
```

## Web server khới động tại địa chỉ: [http://localhost:8080/](http://localhost:8080/ "http://localhost:8080/")

## Truy cập vào địa chỉ: [http://localhost:8080/](http://localhost:8080/ "http://localhost:8080/"). Web server trả về form upload ảnh

![Form upload ảnh](home-page.png?raw=true "Form upload ảnh")

## Sau khi chọn ảnh xong và click Submit, client gửi một POST request lên địa chỉ [http://localhost:8080/upload-image](http://localhost:8080/upload-image "http://localhost:8080/upload-image").

## Server khi nhận được POST request sẽ tiến hành đọc dữ liệu file gửi lên và lưu vào thư mục public. Nếu lưu thành công sẽ redirect sang trang mới với ảnh vừa được upload lên

![Upload ảnh thành công](upload-success.png?raw=true "Upload ảnh thành công")