// Создадим пакет printer с функцией, которая генерирует приветственное сообщение на основе имени человека.
// Создайте папку для проекта и добавьте файл main.go:

// package printer

// import "fmt"

// func PrintHello (name string) string {
// 	return fmt.Sprintf("Hello, %s!", name)
// }

package printer

import "fmt"

var names = make(map[string]string)

func PrintHello(name string) string {
	names[name] = name
	return fmt.Sprintf("Hello, %s!", name)
}

//go test -v -count=3
