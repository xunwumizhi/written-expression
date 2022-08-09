package main

type Book struct {
	Name string
	Age  string
}

func main() {
	var book Book
	book = Book{}

	fmt.Println(book)
}
