package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	even := make(chan int) // unbuffered
	odd := make(chan int)  // unbuffered
	done := make(chan bool)
	go produceEven(even)
	go produceOdd(odd)
	go timeout(done)
	for {
		select {
		case e := <-even:
			fmt.Printf("Even: %d\n", e)
		case o := <-odd:
			fmt.Printf("Odd: %d\n", o)
		case <-done:
			fmt.Println("Done!")
			return
		default:
			//fmt.Println("Nothing received.")
		}
	}
}

func produceEven(nums chan int) {
	for {
		nums <- rand.Intn(100) * 2
	}
}

func produceOdd(nums chan int) {
	for {
		nums <- rand.Intn(100)*2 + 1
	}
}

func timeout(done chan bool) {
	time.Sleep(1 * time.Second)
	done <- true
}
