// Помогите Гоше найти дубликаты. Для этого напишите функцию CountingSort(contacts []string) map[string]int,
// которая принимает слайс строк и возвращает мапу, где ключ - это элемент слайса,
// а значение - количество раз, сколько элемент встречается в слайсе.

package main

func CountingSort(contacts []string) map[string]int {

	newMap := make(map[string]int)

	for _, num := range contacts {
		// Увеличиваем счетчик для каждого элемента
		newMap[num]++
	}
	return newMap
}
