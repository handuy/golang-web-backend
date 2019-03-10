# Chaỵ lệnh

```go
go run main.go
```

## Web server khới động tại địa chỉ: [http://localhost:8080/](http://localhost:8080/ "http://localhost:8080/")

## Truy cập vào đường dẫn /new: [http://localhost:8080/new](http://localhost:8080/new "http://localhost:8080/new"). Server trả về: Đây là router /new

![Đường dẫn new](router-new.png?raw=true "Đường dẫn new")

## Truy cập vào đường dẫn /old: [http://localhost:8080/old](http://localhost:8080/old "http://localhost:8080/old"). Server trả về: "Đây là router /old" và sau đó điều hướng request sang đường dẫn /new

![Đường dẫn old](router-old.png?raw=true "Đường dẫn old")

## Truy cập vào đường dẫn /google: [http://localhost:8080/google](http://localhost:8080/google "http://localhost:8080/google"). Server điều hướng request sang trang chủ Google