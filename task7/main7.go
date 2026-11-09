package main

import (
	"fmt"
	"sync"
)

// Напишите программу на Go, которая сливает N каналов в один.
//
// * Напишите unit тесты к созданным функциям
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer close(ch1)
		ch1 <- 1
		ch1 <- 2
	}()

	go func() {
		defer close(ch2)
		ch2 <- 3
		ch2 <- 4
	}()

	go func() {
		defer close(ch3)
		ch3 <- 5
		ch3 <- 6
	}()

	for n := range mergeChanns(ch1, ch2, ch3) {
		fmt.Println(n)
	}
}

func mergeChanns(channs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range channs {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for n := range ch {
				out <- n
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
