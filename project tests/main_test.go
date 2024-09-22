// Затем мы пишем простой модульный тест для него.
// Добавим тесты в файл main_test.go:

package printer

import "testing"

func TestPrintHello(t *testing.T) {
	got := PrintHello("Igor")
	expected := "Hello, Igor!"

	if got != expected {
		t.Fatalf(`PrintHello("Igor") = %q, want %q`, got, expected)
	}
}

// go mod init module3_1-lesson
// go test
// go test -v

// func TestPrintHello(t *testing.T) {
// 	t.Fail()
// }

// Команда go clean удаляет временные файлы, которые были созданы во время сборки проекта.
// Она также удаляет кэш результатов тестов. Давайте попробуем запустить тесты после очистки кэша:

// Флаг -count - несколько раз запускать

// Теперь давайте изменим тест так, чтобы он проверял, что в map есть имя Igor:
// func TestPrintHello(t *testing.T) {
// 	got := PrintHello("Igor")
// 	expected := "Hello, Igor!"

// 	if got != expected {
// 		t.Fatalf(`PrintHello("Igor") = %q, want %q`, got, expected)
// 	}

// 	for name := range names {
// 		if name != "Igor" {
// 			t.Fatalf(`got %q, want %q`, name, "Igor")
// 		} else {
// 			break
// 		}
// 	}
// }

//тест упадет: (или нет)

// func TestPrintHello(t *testing.T) {
// 	names["Petr"] = "Petr"
// 	got := PrintHello("Igor")
// 	expected := "Hello, Igor!"

// 	if got != expected {
// 		t.Fatalf(`PrintHello("Igor") = %q, want %q`, got, expected)
// 	}

// 	for name := range names {
// 		if name != "Igor" {
// 			t.Fatalf(`got %q, want %q`, name, "Igor")
// 		} else {
// 			break
// 		}
// 	}
// }

// go test -v -run=TestPrintHello - запуск отдельных тестов

//go test -v -run=^TestPrintHello$ - если бы еще был тест, например TestPrintHelloIVAN

//Опция ./... позволяет запускать тесты во всех пакетах внутри текущего модуля. Давайте попробуем её в деле:

-v — выводит подробную информацию о тестах
go clean — очищает кэш результатов тестов
-count=1 — запускает тесты нужное количество раз
-run — запускает тесты по имени
./... — запускает тесты во всех пакетах внутри текущего модуля
-cover — выводит информацию о покрытии кода тестами