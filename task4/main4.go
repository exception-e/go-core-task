package main

import "fmt"

//На вход подаются два неупорядоченных слайса строк.
//Например:
//```
//slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
//slice2 := []string{"banana", "date", "fig"}
//```
//Напишите функцию, которая возвращает слайс строк, содержащий элементы, которые есть в первом слайсе, но отсутствуют во втором.
//
//* Напишите unit тесты к созданным функциям

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	fmt.Println(sliceDiff(slice1, slice2))
}

// сохраняет дубли
func sliceDiff(slice1, slice2 []string) []string {
	result := make([]string, 0, len(slice1))
	mapSlice2 := make(map[string]struct{}, len(slice2))

	for _, v := range slice2 {
		mapSlice2[v] = struct{}{}
	}
	for _, v := range slice1 {
		if _, ok := mapSlice2[v]; !ok {
			result = append(result, v)
		}
	}
	return result
}
