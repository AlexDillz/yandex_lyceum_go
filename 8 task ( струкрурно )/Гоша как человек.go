// Создайте структуру Person с полями name, age и address.
// Создайте метод Print для этой структуры,
// который будет выводить информацию о человеке на экран в формате:

// Name: Гоша
// Age: 21
// Address: Ясногорск

package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
}

func (p Person) Print() {
	fmt.Printf("Name: %s\nAge: %d\nAddress: %s\n", p.name, p.age, p.address)
}
