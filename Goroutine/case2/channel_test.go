package main

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	Server := NewServer()
	Server.Start()
	var wg sync.WaitGroup // 1. 声明一个 WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			Server.userch <- fmt.Sprintf("user%d", i)
		}(i)
	}

	// adding a time to wait for the goroutine to finish
	// time.Sleep(time.Second)
	wg.Wait()
	assert.Equal(t, 10, len(Server.users))
}
