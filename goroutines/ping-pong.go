package main

import (
	"fmt"
	"time"
)

func ping(ch chan string) {
	ch <- "ping"
}

func pong(ch chan string) {
	<-ch
	fmt.Println("pong")
	ch <- "pong"
}

func pingPong() {
	ch := make(chan string)

	for i := 0; i < 3; i++ {
		go ping(ch)
		go pong(ch)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Main over")
}