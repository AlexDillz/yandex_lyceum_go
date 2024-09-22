// Создайте интерфейс Animal с методом MakeSound, который будет выводить звук, издаваемый животным.
// Создайте структуры Dog и Cat, которые будут реализовывать этот интерфейс и издавать соответствующие звуки (выводить на экран) – "Гав!" и "Мяу!".

package main

import (
	"fmt"
)

type Animal interface {
	MakeSound()
}

type Dog struct{}

type Cat struct{}

func (d Dog) MakeSound() {
	fmt.Println("Гав!")
}

func (c Cat) MakeSound() {
	fmt.Println("Мяу!")
}

func main() {
	// Создаем экземпляры Dog и Cat
	var animal Animal

	dog := Dog{}
	cat := Cat{}

	// Присваиваем собаку интерфейсу Animal и выводим звук
	animal = dog
	animal.MakeSound()

	// Присваиваем кошку интерфейсу Animal и выводим звук
	animal = cat
	animal.MakeSound()
}
