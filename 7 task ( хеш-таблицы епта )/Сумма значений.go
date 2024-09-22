// Напишите функцию SumOfValuesInMap(m map[int]int) int, которая возвращает сумму значений в мапе.

package main

func SumOfValuesInMap(m map[int]int) int {

	var sum int

	for _, value := range m {
		sum += value
	}
	return sum

}
