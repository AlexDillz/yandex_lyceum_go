// Создайте интерфейс Logger с методом Log, который будет записывать сообщение в журнал.
// Создайте пользовательский тип LogLevel типа string, сделайте константные значения типа LogLevel со значениями Error и Info.
// Создайте структуру Log с полем LogLevel.
// Реализуйте метод Log c параметром типа string (текст ошибки), который в зависимости от текущего LogLevel выводит сообщение "ERROR: *текст ошибки*" или "INFO: *текст ошибки*".

package main

import (
	"fmt"
)

type Logger interface {
	Log(message string)
}

type LogLevel string

const (
	Error LogLevel = "ERROR"
	Info  LogLevel = "INFO"
)

type Log struct {
	Level LogLevel
}

func (l Log) Log(message string) {
	if l.Level == Error {
		fmt.Println("ERROR:", message)
	} else if l.Level == Info {
		fmt.Println("INFO:", message)
	}
}
