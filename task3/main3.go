package main

import "fmt"

//Реализуйте структуру данных StringIntMap, которая будет использоваться для хранения пар "строка - целое число". Ваша
//структура должна поддерживать следующие операции:
//
//1. Добавление элемента: Метод Add(key string, value int), который добавляет новую пару "ключ-значение" в карту.
//
//2. Удаление элемента: Метод Remove(key string), который удаляет элемент по ключу из карты.
//
//3. Копирование карты: Метод Copy() map[string]int, который возвращает новую карту, содержащую все элементы текущей карты.
//
//4. Проверка наличия ключа: Метод Exists(key string) bool, который проверяет, существует ли ключ в карте.
//
//5. Получение значения: Метод Get(key string) (int, bool), который возвращает значение по ключу и булевый флаг, указывающий на успешность операции.
//
//* Напишите unit тесты к созданным функциям

func main() {
	stringIntMap := New()
	fmt.Println("Put \"first\":1 to map")
	stringIntMap.Add("first", 1)
	fmt.Println(stringIntMap)
}

type StringIntMap struct {
	data map[string]int
}

func New() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

func (m *StringIntMap) Add(key string, value int) {
	if m.data == nil {
		m.data = make(map[string]int)
	}
	m.data[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

func (m *StringIntMap) Copy() map[string]int {
	result := make(map[string]int, len(m.data))
	for k, v := range m.data {
		result[k] = v
	}
	return result
}

func (m *StringIntMap) Exists(key string) bool {
	_, ok := m.data[key]
	return ok
}

func (m *StringIntMap) Get(key string) (int, bool) {
	v, ok := m.data[key]
	return v, ok
}
