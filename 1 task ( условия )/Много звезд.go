package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scanln(&number)

	switch {
	case number < 10:
		fmt.Println("Число меньше 10")
	case number >= 10 && number < 100:
		fmt.Println("Число меньше 100")
	case number >= 100 && number < 1000:
		fmt.Println("Число меньше 1000")
	case number >= 1000:
		fmt.Println("Число больше или равно 1000")
	case number < 0:
		fmt.Println("Некорректный ввод")
	}
}
