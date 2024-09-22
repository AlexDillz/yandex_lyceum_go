// Напишите функцию TimeDifference(start, end time.Time) time.Duration, которая вычисляет разницу во времени между двумя заданными моментами времени.
// Функция должна принимать на вход два значения time.Time и возвращать продолжительность между ними.

package main

import "time"

func TimeDifference(start, end time.Time) time.Duration {

	return end.Sub(start)
}
