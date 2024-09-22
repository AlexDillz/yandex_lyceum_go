package main

import (
	"fmt"
)

func main() {
	var first_number, second_number int
	fmt.Scanln(&first_number, &second_number)

	switch {
	case first_number > 0 && second_number > 0:
		fmt.Println("Оба числа положительные")
	case first_number < 0 && second_number < 0:
		fmt.Println("Оба числа отрицательные")
	case first_number > 0 && second_number < 0 || first_number < 0 && second_number > 0:
		fmt.Println("Одно число положительное, а другое отрицательное")
	default:
		fmt.Println("Одно из чисел равно нулю")
	}
}
