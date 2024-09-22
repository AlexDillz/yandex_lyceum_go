// Напишите функцию ParseStringToTime(dateString, format string) (time.Time, error),
// которая анализирует строку в определенном формате и преобразует ее в значение time.Time.
//  Функция должна принимать строку и строку формата, возвращая проанализированное значение time.Time.

package main

import "time"

func ParseStringToTime(dateString, format string) (time.Time, error) {
	parsedTime, err := time.Parse(format, dateString)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}
