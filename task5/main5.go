package main

import "fmt"

//На вход подаются два неупорядоченных слайса int любой длины.
//Например:
//```
//a := []int{65, 3, 58, 678, 64}
//b := []int{64, 2, 3, 43}
//```
//Напишите функцию, которая проверяет, есть ли пересечения значений между двумя слайсами и возвращает:
//- bool значение есть ли хотя бы одно пересечение в значениях входных срезов
//- срез []int с пересеченными значениями (если таких значений нет, то возвращать пустой срез).
//Т. е. если взать слайсы a и b из премере, то вернет ```true, []int{64, 3}```.
//
//* Напишите unit тесты к созданным функциям

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	fmt.Println(intersection(a, b))
}

func intersection(slice1, slice2 []int) (bool, []int) {
	result := make([]int, 0, len(slice1))
	mapSlice1 := make(map[int]struct{}, len(slice1))
	mapSlice2Dup := make(map[int]struct{}, len(slice2))

	for _, v := range slice1 {
		mapSlice1[v] = struct{}{}
	}

	for _, v := range slice2 {
		_, ok := mapSlice1[v]
		_, dup := mapSlice2Dup[v]
		if ok && !dup {
			result = append(result, v)
			mapSlice2Dup[v] = struct{}{}
		}
	}

	return len(result) > 0, result
}
