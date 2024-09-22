package main

import (
	"fmt"
)

func main() {
	var first_number, second_number, third_number int

	// Считываем три числа и проверяем успешность ввода
	if _, err := fmt.Scan(&first_number, &second_number, &third_number); err != nil {
		fmt.Println("Некорректный ввод")
		return
	}

	// Используем switch для сравнения чисел
	switch {
	case first_number == second_number && second_number == third_number:
		fmt.Println("Все числа равны")
	case first_number == second_number || second_number == third_number || first_number == third_number:
		fmt.Println("Два числа равны")
	default:
		fmt.Println("Все числа разные")
	}
}
