package main 
import "fmt"
type Books struct{
	title string
	id int
}
func main() { 
	var book1 Books
	var book2 Books
	book1.title="samit_Book"
	book1.id=10
	book2.title="samit_Book1"
	book2.id=20
	fmt.Printf("book1 title %s\n",book1.title)
	fmt.Printf("book2 title %s\n",book2.title)
	fmt.Printf("book1 id %d\n",book1.id)
	fmt.Printf("book2 id %d\n",book2.id)
} 
