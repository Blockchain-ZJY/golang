package main

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("url :%s ,elaspe: %v", r.URL, time.Since(now))

	}
	// 将实例化函数赋值给http.HandlerFunc，返回一个满足接口要求（ServeHTTP）的实例
	return http.HandlerFunc(fn)
}

// 装饰器模式，用于对多个常用共性的方法添加一些装饰，比如日志，权限，缓存等
func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func Hi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi, World!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/Hello", Hello)
	mux.HandleFunc("/Hi", Hi)

	server := &http.Server{
		Addr:    ":8080",
		Handler: Logger(mux),
	}

	server.ListenAndServe()
}
