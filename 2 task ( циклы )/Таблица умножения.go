package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scanln(&number)

	for i := 1; i <= 10; i++ {
		res := number * i
		fmt.Println(number, "*", i, "=", res)
	}
}
