// Создайте структуру Student с полями name (строка),
// solvedProblems — количество решённых задач (целое число), scoreForOneTask— количество баллов за одну задачу;
// будем считать, что все задачи дают одинаковые баллы (число с плавающей точкой) и passingScore —
// проходной балл в следующий модуль (число с плавающей точкой).

// Создайте метод IsExcellentStudent для этой структуры, который возвращает true,
// если ученик проходит по баллам в следующий модуль и false в ином случае.

package main

type Student struct {
	solvedProblems  int
	scoreForOneTask float64
	passingScore    float64
}

func (s Student) IsExcellentStudent() bool {

	totalScore := float64(s.solvedProblems) * s.scoreForOneTask

	return totalScore >= s.passingScore
}
