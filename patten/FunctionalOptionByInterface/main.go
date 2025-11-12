package main

import (
	"fmt"
	"time"
)

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
}

// Option 接口，定义需要实现的apply方法
type Option interface {
	apply(server *Server)
}

// ProtoOption 实现Option接口
type ProtoOption string

func (p ProtoOption) apply(s *Server) {
	s.Protocol = string(p)
}

func WithProtocol(proto string) Option {
	return ProtoOption(proto)
}

// TimeoutOption 实现Option接口
type TimeoutOption time.Duration

func (t TimeoutOption) apply(s *Server) {
	s.Timeout = time.Duration(t)
}

func WithTimeout(timeout time.Duration) Option {
	return TimeoutOption(timeout)
}

// MaxConnOption 实现Option接口
type MaxConnOption int

func (m MaxConnOption) apply(s *Server) {
	s.MaxConns = int(m)
}

func WithMaxConn(maxConn int) Option {
	return MaxConnOption(maxConn)
}

func NewServer(addr string, port int, opts ...Option) *Server {
	//创建Server，并填写可选项的默认值
	s := &Server{
		Addr:     "127.0.0.1",
		Port:     8080,
		Protocol: "udp",
		Timeout:  time.Second * 10,
		MaxConns: 10,
	}

	for _, opt := range opts {
		opt.apply(s)
	}

	return s
}

/****************************************************/
func main() {
	s := NewServer("192.168.101.84", 1234, WithProtocol("tcp"), WithTimeout(time.Minute), WithMaxConn(100))
	fmt.Printf("server:%v\n", s)
}
