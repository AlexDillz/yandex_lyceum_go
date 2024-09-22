// Напишите программу, которая запрашивает у пользователя число и выводит на экран
// все числа от 1 до этого числа, которые делятся на 3 с помощью цикла for.

package main

import "fmt"

func main() {
	var number int
	fmt.Scanln(&number)

	for i := 1; i <= number; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}

}
