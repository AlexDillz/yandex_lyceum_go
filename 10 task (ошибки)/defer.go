func main() {
	defer fmt.Println("World")
	fmt.Println("Hello")
}

// Hello
// World

// Заметим, что несколько defer исполняются в обратном порядке:

func main() {
	defer fmt.Println("World")
	defer fmt.Println("Hello")
}

// Вывод:
// Hello
// World