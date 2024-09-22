// Напишите программу, которая запрашивает у пользователя число и выводит на экран сумму всех нечётных чисел от 1 до этого числа
// с помощью цикла for.
// Если число отрицательное, программа должна выводить сообщение Некорректный ввод.

// package main

// import "fmt"

// func main() {
// 	var number int
// 	fmt.Scanln(&number)
// 	res := 1

// 	for i := 1; i <= number; i++ {
// 		if i == number+1 {
// 			break
// 		}
// 		res *= i
// 	}
// 	fmt.Print(res)
// }

package main

import "fmt"

func main() {

	var number int
	fmt.Scanln(&number)
	res := 0

	if number >= 0 {
		for i := 1; i <= number; i++ {
			if i%2 == 0 {
				continue
			}
			res += i
		}
		fmt.Println(res)
	}
	if number <= 0 {
		fmt.Println("Некорректный ввод")
	}
}
