// Функция Sum(a, b int) int (пакет main) возвращает результат суммирования чисел a и b.
// Напишите тест TestSum(t *testing.T) для проверки корректности работы.

package main

import "testing"

func TestSum(t *testing.T) {
	if Sum(5, 5) != 10 {
		t.Fatalf("ERROR")
	}
}
