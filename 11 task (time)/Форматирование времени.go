// Создайте функцию FormatTimeToString(timestamp time.Time, format string) string,
// которая форматирует заданное значение time.Time как строку в определенном формате.
// Функция должна принимать значение time.Time и строку формата и возвращать форматированную строку.

package main

import "time"

func FormatTimeToString(timestamp time.Time, format string) string {

	return timestamp.Format(format)
}
