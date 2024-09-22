package main

import (
	"fmt"
)

func main() {
	var first_number, second_number float32
	fmt.Scanln(&first_number, &second_number)

	if first_number > second_number {
		fmt.Println("Первое число больше второго")
	} else if first_number < second_number {
		fmt.Println("Второе число больше первого")
	} else {
		fmt.Println("Числа равны")
	}
}
