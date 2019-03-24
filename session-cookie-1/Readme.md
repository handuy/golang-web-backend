# Chạy lệnh

```go
go run main.go
```

## Web server khới động tại địa chỉ: [http://localhost:8080/](http://localhost:8080/ "http://localhost:8080/")

## Truy cập vào địa chỉ: [http://localhost:8080/secret](http://localhost:8080/secret "http://localhost:8080/secret")

![Chưa đăng nhập](not_log_in.png?raw=true "Chưa đăng nhập")

## Truy cập vào địa chỉ: [http://localhost:8080/login](http://localhost:8080/login "http://localhost:8080/secret")

![Đăng nhập thành công](login.png?raw=true "Đăng nhập thành công")

## Quay lại truy cập vào địa chỉ [http://localhost:8080/secret](http://localhost:8080/secret "http://localhost:8080/secret")

![Secret message](secret.png?raw=true "Secret message")

## Restart server, sau đó truy cập lại vào địa chỉ [http://localhost:8080/secret](http://localhost:8080/secret "http://localhost:8080/secret")

![Secret message](secret.png?raw=true "Secret message")

## Gọi vào đường dẫn [http://localhost:8080/logout](http://localhost:8080/logout "http://localhost:8080/logout")

![Logout](logout.png?raw=true "Logout")

## Truy cập lại vào đường dẫn [http://localhost:8080/secret](http://localhost:8080/secret "http://localhost:8080/secret")

![Chưa đăng nhập](not_log_in.png?raw=true "Chưa đăng nhập")

------------

## Giải thích code

## Khi khởi động, server sẽ tạo một "cửa hàng Cookie" (Cookie Store). Đây là nơi lưu trữ thông tin về session thông qua secured cookie

```go
var (
    key = []byte("my-super-secret-key")
    store = sessions.NewCookieStore(key)
)
```

## Thông tin về session được lưu dưới dạng một struct

```go
type Session struct {
    // ID của session được sinh ra bởi cookie store
    ID string
    // Values lưu thông tin về dữ liệu của user dưới dạng map
    Values  map[interface{}]interface{}
    Options *Options
    // IsNew check xem session mới được tạo ra (true) hay là session cũ (false)
    IsNew   bool
    store   Store
    // name lưu tên của session
    name    string
}
```

## Khi gọi vào đường dẫn /secret, server sẽ gọi vào hàm secret

```go
func secret(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "cookie-name")

    if auth, ok := session.Values["authenticated"].(bool); !auth || !ok {
        http.Error(w, "Bạn chưa đăng nhập", http.StatusUnauthorized)
        return
    }

    fmt.Fprintln(w, "<h1>This is my secret message</h1>")
}
```

## Server sẽ tìm trong cookie store xem có session nào có name = "cookie-name" hay không, nếu không có thì cookie store sẽ tạo session mới. Ở lần đầu tiên chúng ta gọi vào /secret (lúc này chưa đăng nhập), cookie store sẽ không tìm thấy session nào có tên là "cookie-name", do đó sẽ tạo một session mới. Và vì đây là session mới, thế nên trường Values của session sẽ không thể có giá trị ứng với key "Authenticated" --> server trả về lỗi "Bạn chưa đăng nhập"

## Khi gọi vào đường dẫn /login, server sẽ gọi vào hàm login

```go
func login(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "cookie-name")

    session.Values["authenticated"] = true
    session.Save(r,w)
    fmt.Fprintln(w, "<h1>Đăng nhập thành công</h1>")
}
```

## Server sẽ làm những việc sau:

1. Tạo một session mới có name="cookie-name" và IsNew=true (nếu là lần đầu tiên login) / Trả về session có name=cookie-name và IsNew=false
2. Set giá trị cho key "authenticated" = true
3. Lưu vào session
4. Trả về dòng chữ "Đăng nhập thành công" trên browser

## Đến lúc này, khi gọi lại vào đường dẫn /secret, trong cookie store đã có session với name="cookie-name" và authenticated=true --> server trả về dòng chữ "This is my secret message" cho browser

## Ngay cả khi restart server và truy cập lại vào /secret, server vẫn trả về dòng chữ "This is my secret message". Đó là bởi browser vẫn lưu lại cookie được set từ lần đăng nhập trước đó

![Cookie](cookie.png?raw=true "Cookie")

## Khi truy cập vào /logout, logic xử lý cũng tương tự như /login, chỉ có điều server sẽ set key authenticated=false