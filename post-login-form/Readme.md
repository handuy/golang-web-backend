# Chaỵ lệnh

```go
go run main.go
```

## Web server khới động tại địa chỉ: [http://localhost:8080/](http://localhost:8080/ "http://localhost:8080/")

## Truy cập vào địa chỉ: [http://localhost:8080/](http://localhost:8080/ "http://localhost:8080/"). Web server trả về 2 đường link: Đăng kí và đăng nhập

![Trang chủ](home-page.png?raw=true "Trang chủ")

## Khi click Đăng kí, client gửi GET request lên địa chỉ [http://localhost:8080/signup](http://localhost:8080/signup "http://localhost:8080/signup"). Server trả về form đăng kí

![Form đăng kí](signup-get.png?raw=true "Form đăng kí")

## Sau khi điền username và password, click Submit. Client sẽ gửi POST request lên địa chỉ [http://localhost:8080/signup](http://localhost:8080/signup "http://localhost:8080/signup"). Server sẽ đọc dữ liệu form gửi lên và lưu vào trong slice userList. Sau đó, server sẽ redirect sang /signup-success

![Đăng kí thành công](signup-success.png?raw=true "Đăng kí thành công")

## Sau khi đã có username và password, click vào link Đăng nhập. Client sẽ gửi GET request lên địa chỉ [http://localhost:8080/login](http://localhost:8080/login "http://localhost:8080/login"). Server sẽ trả về form login

![Đăng nhập](login.png?raw=true "Đăng nhập")

## Sau khi điền username và password, click Submit. Client sẽ gửi POST request lên địa chỉ [http://localhost:8080/login](http://localhost:8080/login "http://localhost:8080/login"). Server sẽ kiểm tra xem username và password gửi lên có tồn tại trong slice userList không. Nếu có thì sẽ redirect sang /login-success. Nếu không có thì redirect sang /login-fail

![Đăng nhập thành công](login-success.png?raw=true "Đăng nhập thành công")