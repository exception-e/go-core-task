package main

import "fmt"

//Сделать кастомную waitGroup на семафоре, не используя sync.WaitGroup.
//
//* Напишите unit тесты к созданным функциям

func main() {
	const n = 10
	wg := New(n)

	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println("goroutine", i)
		}(i)
	}

	wg.Wait()
	fmt.Println("all done")
}

type MyWaitGroup struct {
	sem chan struct{}
	n   int
}

func New(n int) *MyWaitGroup {
	return &MyWaitGroup{
		sem: make(chan struct{}, n),
		n:   n,
	}
}

func (wg *MyWaitGroup) Done() {
	wg.sem <- struct{}{}
}

func (wg *MyWaitGroup) Wait() {
	for i := 0; i < wg.n; i++ {
		<-wg.sem
	}
}
