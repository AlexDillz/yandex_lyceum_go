// Создайте интерфейс Shape с методом Area, который будет возвращать площадь фигуры.
// Создайте структуры Circle и Rectangle, которые будут реализовывать этот интерфейс и рассчитывать площадь этих фигур.

package main

import "math"

// Интерфейс Shape с методом Area
type Shape interface {
	Area() float64
}

// Структура Circle с полем Radius
type Circle struct {
	Radius float64
}

// Структура Rectangle с полями Width и Height
type Rectangle struct {
	Width  float64
	Height float64
}

// Реализация метода Area для Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Реализация метода Area для Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
