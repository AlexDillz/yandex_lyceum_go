// Создайте структуру Task:
// 1. summary - короткий заголовок (тип string)
// 2. description - подробное описание (тип string)
// 3. deadline - дедлайн для задачи (тип time.Time (стандартный пакет time))
// 4. priority - приоритет задачи (тип int)

// Для этой структуры реализуйте метод IsOverdue, которая возвращает true, если дедлайн задачи уже прошел и false в ином случае.
// Подсказка: воспользуйтесь функцией time.Now().

// Ещё запрограммируйте метод IsTopPriority. Функция возвращает true, если приоритет 4 или 5, и false, если меньше.

package main

import "time"

type Task struct {
	summary     string
	description string
	deadline    time.Time
	priority    int
}

func (t Task) IsOverdue() bool {

	if time.Now().After(t.deadline) {
		return true
	} else {
		return false
	}

}

func (t Task) IsTopPriority() bool {
	if t.priority == 4 || t.priority == 5 {
		return true
	} else {
		return false
	}
}
