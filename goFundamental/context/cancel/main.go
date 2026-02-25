package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func calculatePos() (x, y int) {
	fmt.Println("calculating the position of the driver")
	x = rand.Intn(100)
	y = rand.Intn(100)
	return
}
func sendRes(x, y int) {
	fmt.Println("the position of driver is", x, ",", y)
}
func Perform(ctx context.Context) {
	for {
		x, y := calculatePos()
		sendRes(x, y)
		select {
		case <-ctx.Done():
			fmt.Println("User stop the svs")
			return
		case <-time.After(time.Second):
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	go Perform(ctx)
	time.Sleep(time.Second * 3)
	cancel()
	time.Sleep(time.Second * 3)
}
