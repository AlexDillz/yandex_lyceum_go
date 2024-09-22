// Напишите функцию SumOfArray(array [6]int) int которая получает массив из шести целочисленных элементов
// и возвращает сумму всех элементов массива.
// Примечания
// Например, если передать функции SumOfArray(array [6]int) int слайс [1 2 3 4 5 6], то она должна вернуть число 21.

package main

func SumOfArray(array [6]int) int {

	var sum int

	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return sum
}
