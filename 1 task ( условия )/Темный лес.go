package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scanln(&number)

	switch {
	case number > 0 && number%2 == 0:
		fmt.Println("Число положительное и четноe")
	case number > 0 && number%2 != 0:
		fmt.Println("Число положительное и нечетное")
	case number < 0 && number%2 == 0:
		fmt.Println("Число отрицательное и четное")
	case number < 0 && number%2 != 0:
		fmt.Println("Число отрицательное и нечетное")
	default:
		fmt.Println("Число равно нулю")
	}
}
