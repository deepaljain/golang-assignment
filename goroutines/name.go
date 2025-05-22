package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printName(name string) {
	time.Sleep(time.Duration(rand.Intn(4)) * time.Second)
	fmt.Println("Hello,", name)
}

func printNames() {
	go printName("John")
	go printName("Doe")
	go printName("Alice")
	go printName("Bob")
	go printName("Charlie")
	time.Sleep(5 * time.Second)
	fmt.Println("Main over")
}