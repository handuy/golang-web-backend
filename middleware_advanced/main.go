package main

import (
	"fmt"
	"log"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func LogPath() Middleware {
	log.Println("Chạy hàm LogPath")
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println("Tên đường dẫn:", r.URL.Path)
			fmt.Fprintln(w, "<h1>Chạy hàm LogPath</h1>")
			f(w, r)
		}
	}
}

func LogMethod() Middleware {
	log.Println("Chạy hàm LogMethod")
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println("Tên method:", r.Method)
			fmt.Fprintln(w, "<h1>Chạy hàm LogMethod</h1>")
			f(w, r)
		}
	}
}

func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	log.Println("Chạy hàm Chain")
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello World</h1>")
}

func main() {
	http.HandleFunc( "/", Chain( Hello, LogPath(), LogMethod() ) )
	http.ListenAndServe(":8080", nil)
}
