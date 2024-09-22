// Напишите структура Note (сущность заметок, у которых в отличие от задач нет чётких дедлайнов и приоритета):
// 1. title - заголовок (тип string)
// 2. text - текст заметок (тип string)

// Создайте структуру ToDoList с такими полями:
// 1. name - название списка (тип string)
// 2. tasks - список дел на сегодня (тип слайс структур Task (из предыдущего задания))
// 3. notes - список дополнительных заметок (тип слайс структур Note)

// Для этой структуры реализуйте методы TasksCount и NotesCount, которые возвращают общее количество задач и заметок соответственно.
// Также реализуйте метод CountTopPrioritiesTasks, который возвращает количество приоритетных задач.
// А также метод CountOverdueTasks, который возвращает количество просроченных задач.
// Сама структура Task и все её методы из предыдущего задания также должны быть реализованы в этом.

package main

import (
	"time"
)

type Note struct {
	title string
	text  string
}

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

type ToDoList struct {
	name  string
	tasks []Task
	notes []Note
}

func (t ToDoList) TasksCount() int {

	return len(t.tasks)
}

func (t ToDoList) NotesCount() int {

	return len(t.notes)
}

func (t ToDoList) CountTopPrioritiesTasks() int {
	count := 0
	for _, task := range t.tasks {
		if task.IsTopPriority() {
			count++
		}
	}
	return count
}

func (t ToDoList) CountOverdueTasks() int {
	count := 0
	for _, task := range t.tasks {
		if task.IsOverdue() {
			count++
		}
	}
	return count
}
