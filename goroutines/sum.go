package main

import "fmt"

func sum(numbers []int, ch chan int) {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	ch <- sum
}

func calculateSum() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	numbers := []int{1, 2, 3, 4, 5}

	mid := len(numbers) / 2
	go sum(numbers[:mid], ch1)
	go sum(numbers[mid:], ch2)

	result := <-ch1 + <-ch2
	fmt.Println("Sum of numbers:", result)
}