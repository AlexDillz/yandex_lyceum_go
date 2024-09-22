// Напишите программу, которая запрашивает у пользователя число — количество букв в алфавите,
// а потом выводит на экран факториал этого числа (количество перестановок букв) с помощью цикла for.

package main

import "fmt"

func main() {
	var number int
	fmt.Scanln(&number)
	res := 1

	for i := 1; i <= number; i++ {
		if i == number+1 {
			break
		}
		res *= i
	}
	fmt.Print(res)
}
