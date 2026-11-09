package main

import (
	"context"
	"fmt"
	"math/rand/v2"
)

//Напишите генератор случайных чисел используя небуфферизированный канал.
//
//* Напишите unit тесты к созданным функциям

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	randCh := randGen(ctx)
	for v := range randCh {
		fmt.Println(v)
	}
}

func randGen(ctx context.Context) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for {
			randN := rand.IntN(1000)
			select {
			case ch <- randN:
			case <-ctx.Done():
				return
			}
		}
	}()
	return ch
}
