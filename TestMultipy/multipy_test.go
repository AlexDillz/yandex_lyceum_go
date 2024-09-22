// Функция Multiply(a, b int) int (пакет main) возвращает произведение двух переданных чисел.
// Напишите тест TestMultiply(t *testing.T) для проверки корректности работы.
package main

import "testing"

func TestMultiply(t *testing.T) {
	if Multiply(2, 5) != 10 {
		t.Fatalf("ERROR")
	}
	if Multiply(2, 0) != 0 {
		t.Fatalf("ERROR")
	}
}

// go mod init module3_1-lesson
// -count=1 — запускает тесты нужное количество раз
