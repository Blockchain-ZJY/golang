package main

import (
	"fmt"
)

type Server struct {
	users  map[string]string
	userch chan string
}

func NewServer() *Server {
	return &Server{
		users:  make(map[string]string),
		userch: make(chan string),
	}
}

func (s *Server) Start() {
	go s.listener()
}

// if the userch has the buffer, the listener will not block,
// if not, the listener will stay for other goroutines to add users until all the
// goroutines are finished
func (s *Server) listener() {
	for {
		user := <-s.userch
		s.users[user] = user
		fmt.Println("user:", user, "joined")
	}
}

func main() {
	Server := NewServer()
	Server.Start()

	for i := 1; i <= 10; i++ {
		go func(i int) {
			Server.userch <- fmt.Sprintf("user %d", i)
		}(i)
	}

	// adding a time to wait for the goroutine to finish
	// time.Sleep(time.Second)

}
