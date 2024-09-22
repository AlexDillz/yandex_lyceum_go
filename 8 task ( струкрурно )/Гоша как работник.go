// Создайте структуру Employee с полями name (string), position (string), salary (float64) и bonus (float64).
// Создайте метод CalculateTotalSalary для этой структуры, который будет выводить общую зарплату работника
// (salary + bonus) по следующему формату:

// Employee: Arseniy

// Position: backend developer

// Total Salary: 1000.02

// Обратите внимание, что Total Salary нужно округлять до 2 знаков после запятой.

package main

import "fmt"

type Employee struct {
	name     string
	position string
	salary   float64
	bonus    float64
}

func (e Employee) CalculateTotalSalary() {
	fmt.Printf("Employee: %s\nPosition: %s\nTotal Salary: %.2f\n", e.name, e.position, e.salary+e.bonus)
}
