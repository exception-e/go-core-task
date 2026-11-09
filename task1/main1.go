package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"strings"
)

//Напишите программу на Go, которая:
//
//1. Создает несколько переменных различных типов данных:
//```
//int (три числа в десятичной, восьмеричной и шеснадцатиричной системах)
//float64
//string
//bool
//complex64
//```
//2. Определяет тип каждой переменной и выводит его на экран.
//3. Преобразует все переменные в строковый тип и объединяет их в одну строку.
//4. Преобразовать эту строку в срез рун.
//5. Захэшировать этот срез рун SHA256, добавив в середину соль "go-2024" и вывести результат.
//
//* Напишите unit тесты к созданным функциям
//
//Напишите main функцию, в которой протестируете весь вышеописанный функционал. Выведите результаты на экран.
//
//Входные числа из пункта 1 могут быть:
//```
//var numDecimal int = 42           // Десятичная система
//var numOctal int = 052            // Восьмеричная система
//var numHexadecimal int = 0x2A     // Шестнадцатиричная система
//var pi float64 = 3.14             // Тип float64
//var name string = "Golang"         // Тип string
//var isActive bool = true           // Тип bool
//var complexNum complex64 = 1 + 2i  // Тип complex64
//```

func main() {
	//1
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	vars := []any{numDecimal,
		numOctal,
		numHexadecimal,
		pi,
		name,
		isActive,
		complexNum}
	formatMask := []string{"%d", "%#o", "%#x", "%f", "%s", "%t", "%g"}

	w := os.Stdout

	salt := []rune("go-2024")
	//2
	printVars(w, vars, formatMask)
	//3
	str := makeStringOfVars(vars, formatMask)
	fmt.Println(str)
	//4
	runes := []rune(str)
	//5
	fmt.Println(hashRunes(runes, salt))
}

func printVars(w io.Writer, vars []any, formatMask []string) {
	for i, v := range vars {
		fmt.Fprintf(w, "type of "+formatMask[i]+": %T\n", v, v)
	}
}

func makeStringOfVars(vars []any, formatMask []string) string {
	s := strings.Builder{}
	for i, v := range vars {
		s.WriteString(fmt.Sprintf(formatMask[i], v))
	}
	return s.String()
}

func hashRunes(runes []rune, salt []rune) [32]byte {
	mid := len(runes) / 2
	saltedRunes := []byte(string(runes[0:mid]))
	saltedRunes = append(saltedRunes, []byte(string(salt))...)
	saltedRunes = append(saltedRunes, []byte(string(runes[mid:]))...)
	return sha256.Sum256(saltedRunes)
}
