package main

import "fmt"

//Сделать конвейер чисел
//Даны два канала.
//В первый пишутся числа типа uint8. Нужно, чтобы числа читались из первого канала по мере поступления,
//затем эти числа должны преобразовываться в float64 и возводиться в куб и результат записывался во второй канал.
//
//Напишите main функцию, в которой протестируете весь вышеописанный функционал. Выведите результаты на экран.
//
//* Напишите unit тесты к созданным функциям

func main() {
	const n = 10
	in := make(chan uint8)

	go func() {
		defer close(in)
		for i := 1; i <= n; i++ {
			in <- uint8(i)
		}
	}()

	out := pipe(in)

	for v := range out {
		fmt.Printf("%.0f\n", v)
	}
}

func pipe(in <-chan uint8) <-chan float64 {
	out := make(chan float64)
	go func() {
		defer close(out)
		for v := range in {
			f := float64(v)
			out <- f * f * f
		}
	}()
	return out
}
