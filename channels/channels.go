package main

import "fmt"

func main() {

	array := []int{8, 7, 9, -9, 7, 5, 2}

	channel := make(chan int)

	go sum(array[:len(array)/2], channel)
	go sum(array[len(array)/2:], channel)

	x, y := <-channel, <-channel

	fmt.Println(x, y, x+y)
}

func sum(s []int, c chan int) {
	sum := 0

	for _, v := range s {
		sum += v
	}

	c <- sum
}
