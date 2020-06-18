package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	channel := make(chan int)

	go func() {
		for i := 1; i <= 4; i++ {
			go queue(channel, i)
		}
	}()

	for i := 0; i < 100; i++ {
		channel <- i
	}
}

func queue(c chan int, worker int) {

	for i := 0; i <= 100; i++ {
		fmt.Println("Número ", i, "Worker: ", worker)
		time.Sleep(time.Duration(rand.Intn(4)) * time.Second)
	}

}
