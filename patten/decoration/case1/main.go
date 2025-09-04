package main

import (
	"log"
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request)

// 装饰器本身，接受进来一个函数，返回一个本身函数
// 传递的参数如果是函数，需要将函数封装成一个类型
func Logger(handler Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Logger")
		handler(w, r)
	}
}

// 装饰器模式，用于对多个常用共性的方法添加一些装饰，比如日志，权限，缓存等
func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

// 装饰器模式，用于对多个常用共性的方法添加一些装饰，比如日志，权限，缓存等
func Hi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi, World!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/Hello", Logger(Hello))
	mux.HandleFunc("/Hi", Logger(Hi))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
