// Функциональный оператор — импровизированный калькулятор с несколькими встроенными операциями. Его алгоритм прост:
// он принимает операцию и параметры для обработки и выдаёт результат.

// Функции, которые нужно реализовать в функциональном операторе:

// Операция	Сигнатура функции	Описание
// Сложение	Add(a, b float64) float64	Возвращает a + b
// Умножение	Multiply(a, b float64) float64	Возвращает a * b
// Возрастающая последовательность	PrintNumbersAscending(n int)
// Печатает возрастающую последовательность от 1 до (включая) n через пробел

// Как и в предыдущем задании, вам нужно написать программу полностью кроме функции main().

package main

import (
	"fmt"
)

func Add(a, b float64) float64 {
	return a + b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func PrintNumbersAscending(n int) {

	fmt.Scanln(&n) // не обязательно

	for i := 1; i <= n; i++ {
		fmt.Println(i) // fmt.Print(" ") для записи с пробелом
	}
	fmt.Println() // Перенос курсора на новую строку
}

func main() {
	// Пример использования функции Add
	resultAdd := Add(3.5, 2.7)
	fmt.Println("Результат сложения:", resultAdd)

	// Пример использования функции Multiply
	resultMultiply := Multiply(4.0, 5.0)
	fmt.Println("Результат умножения:", resultMultiply)

	// Пример использования функции PrintNumbersAscending
	fmt.Print("Последовательность чисел: ")
	PrintNumbersAscending(5)
}
