package main 
import "fmt"
func swap(a,b int) int{
	var o int
	o=a
	a=b
	b=o
	return o
}
func main(){
	var a int=10
	var b int=20
	fmt.Printf("a=%d b=%d\n",a,b)
	fmt.Println("After swap")
	swap(a,b)
	fmt.Printf("a=%d b=%d",a,b)
}