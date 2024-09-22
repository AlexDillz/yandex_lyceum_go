// Напишите функцию FindIntersection(k1, b1, k2, b2 float64) (float64, float64),
// которая будет принимать на вход четыре вещественных числа k1, b1, k2 и b2 — они представляют коэффициенты уравнений y = kx + b
// для двух разных астрономических объектов. Ваша задача — найти точку, в которой эти астрономические траектории пересекаются.
// Если траектории параллельны, программа должна вывести math.NaN(), math.NaN().

package main

import (
	"math"
)

func FindIntersection(k1, b1, k2, b2 float64) (float64, float64) {

	if k1 == k2 {
		return math.NaN(), math.NaN()
	}

	x := (b2 - b1) / (k1 - k2)
	Equation := k1*x + b1

	return x, Equation

}
