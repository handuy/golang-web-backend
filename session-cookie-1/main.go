package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/sessions"
)

var (
	key = []byte("my-super-secret-key")
	store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	if auth, ok := session.Values["authenticated"].(bool); !auth || !ok {
		http.Error(w, "Bạn chưa đăng nhập", http.StatusUnauthorized)
		return
	}

	fmt.Fprintln(w, "<h1>This is my secret message</h1>")
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	session.Values["authenticated"] = true
	session.Save(r,w)
	fmt.Fprintln(w, "<h1>Đăng nhập thành công</h1>")
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	session.Values["authenticated"] = false
	session.Save(r,w)
	fmt.Fprintln(w, "<h1>Bạn vừa logout</h1>")
}

func main() {
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	http.ListenAndServe(":8080", nil)
}