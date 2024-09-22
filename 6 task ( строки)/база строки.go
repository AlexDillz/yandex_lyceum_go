// strings.ToLower преобразовывает строки в нижний регистр. Это полезно, если вам нужно сравнить две строки без учёта регистра:

func main() {
	name := "АнДРей"
	lower := strings.ToLower(name)
	fmt.Println(lower) // Output: андрей
}

// strings.Join объединяет строки в одну строку с указанным разделителем:

func main() {
	words := []string{"Это", "несколько", "слов"}
	joined := strings.Join(words, " ")
	fmt.Println(joined) // Output: Это несколько слов
}

// Напоследок напишем фразу «Я ЛЮБЛЮ GO»:

func main() {
	text := "Я ЛЮБЛЮ Golang"
	fmt.Println(strings.Replace(text, "Golang", "GO", -1))
	// Output: Я ЛЮБЛЮ GO
}