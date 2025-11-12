package main

import "fmt"

// 实现可选项模式
type Server struct {
	Host     string
	Port     int
	TLS      bool
	MaxConns int
}

// Option 是一个函数类型，用于修改 Server 的配置。
type Option func(*Server)

func NewServer(options ...Option) *Server {
	// 创建一个具有默认配置的服务器实例
	s := &Server{
		Host:     "localhost",
		Port:     8080,
		TLS:      false,
		MaxConns: 100,
	}
	// 应用所有传入的选项来修改服务器配置
	for _, option := range options {
		option(s)
	}
	return s
}

// 定义一些选项函数

func WithHost(host string) Option {
	return func(s *Server) {
		s.Host = host
	}
}

func WithPort(port int) Option {
	return func(s *Server) {
		s.Port = port
	}
}

func WithTLS(tls bool) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}

func WithMaxConns(maxConns int) Option {
	return func(s *Server) {
		s.MaxConns = maxConns
	}
}

func main() {

	server1 := NewServer()
	server2 := NewServer(WithHost("2020"))
	fmt.Println(server1, server2)
}
