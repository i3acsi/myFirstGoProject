package main

import (
	"fmt"
	"math/rand"
)

func one() {
	nums := make(chan int) // unbuffered
	go produceOne(nums)
	consumeOne(nums)

}

func produceOne(nums chan int) {
	nums <- 1
}

func consumeOne(nums chan int) {
	n := <-nums
	fmt.Println(n)
}

func many(i int) {
	nums := make(chan int, 10) // buffered
	go produce(nums, i)
	consume(nums, i)
}

func produce(nums chan int, length int) {
	i := 0
	for i <= length {
		nums <- rand.Intn(200)
		i += 1
	}
	close(nums)
}

func consume(nums chan int, length int) {
	for n := range nums {
		fmt.Println(n)
	}
}

func main() {
	one()
	many(100)
}
