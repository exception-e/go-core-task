package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
)

//1. Создайте слайс целых чисел originalSlice, содержащий 10 произвольных значений, которые генерируются случайным
//образом (при каждом запуске должны получаться новые значения)
//
//2. Напишите функцию sliceExample, которая принимает слайс и возвращает новый слайс, содержащий только четные числа из исходного слайса.
//
//3. Напишите функцию addElements, которая принимает слайс и число. Функция должна добавлять это число в конец слайса и возвращать новый слайс.
//
//4. Напишите функцию copySlice, которая принимает слайс и возвращает его копию. Убедитесь, что изменения в оригинальном слайсе не влияют на его копию.
//
//5. Напишите функцию removeElement, которая принимает слайс и индекс элемента, который нужно удалить. Функция должна возвращать новый слайс без элемента по указанному индексу.
//
//6. Напишите main функцию, в которой протестируете все вышеописанные функции. Выведите результаты на экран.
//
//* Напишите unit тесты к созданным функциям
//
//Примечание.
//В качестве originalSlice можно использовать ```originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}```

func main() {
	originalSlice := make([]int, 10)
	for i := range originalSlice {
		originalSlice[i] = rand.IntN(1000)
	}

	fmt.Println(originalSlice)
	fmt.Println(sliceExample(originalSlice))
	n := rand.IntN(1000)
	fmt.Printf("addElements(%v) → ", n)
	fmt.Println(addElements(originalSlice, n))
	fmt.Println(copySlice(originalSlice))
	fmt.Println(removeElement(originalSlice, 3))
}

func sliceExample(originalSlice []int) []int {
	result := make([]int, 0, len(originalSlice))
	for _, v := range originalSlice {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

func addElements(originalSlice []int, n int) []int {
	result := append([]int(nil), originalSlice...)
	result = append(result, n)
	return result
}
func copySlice(originalSlice []int) []int {
	return slices.Clone(originalSlice)
}
func removeElement(originalSlice []int, n int) ([]int, error) {
	if n < 0 || n >= len(originalSlice) {
		return nil, fmt.Errorf("element out of slice range")
	}
	result := make([]int, 0, len(originalSlice)-1)
	result = append(result, originalSlice[:n]...)
	result = append(result, originalSlice[n+1:]...)
	return result, nil
}
