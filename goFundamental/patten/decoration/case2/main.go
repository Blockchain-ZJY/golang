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
	// 将一个普通的函数转化成http.HandlerFunc类型
	// 任何满足(http.ResponseWriter,*http.Request)签名的函数经过HandlerFunc
	// 都会被认定为是handler的一种，因为HandlerFunc实现了http.Handler接口
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
