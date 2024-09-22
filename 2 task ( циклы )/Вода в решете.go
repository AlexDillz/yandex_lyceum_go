//Напишите программу, которая запрашивает у пользователя число и
//выводит на экран сумму всех чисел от 1 до этого числа с помощью цикла for

package main

import "fmt"

func main() {
	var number int
	fmt.Scanln(&number)
	res := 0

	for i := 1; i <= number; i++ {
		if i == number+1 {
			break
		}
		res += i
	}
	fmt.Print(res)
}
