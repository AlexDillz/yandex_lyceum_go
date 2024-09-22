//бесконечный цикл, который запрашивает ввод пользователя из консоли и проверяет,
//равен ли этот ввод строке "exit". Если пользователь вводит "exit", цикл
//завершается с помощью оператора break:

package main

import (
	"fmt"
)

func main() {
	for {
		var input string
		fmt.Scanln(&input)
		if input == "exit" {
			break
		}
	}
}
