// Напишите функцию NextWorkday(start time.Time) time.Time, которая вычисляет дату следующего рабочего дня (исключая выходные).
// Учитывая дату начала, функция должна возвращать дату следующего рабочего дня.

package main

import (
	"time"
)

func NextWorkday(start time.Time) time.Time {
	// Определяем день недели
	switch start.Weekday() {
	case time.Friday:
		// Следующий рабочий день — понедельник
		return start.AddDate(0, 0, 3)
	case time.Saturday:
		// Следующий рабочий день — понедельник
		return start.AddDate(0, 0, 2)
	case time.Sunday:
		// Следующий рабочий день — понедельник
		return start.AddDate(0, 0, 1)
	default:
		// В обычные дни прибавляем один день
		return start.AddDate(0, 0, 1)
	}
}
