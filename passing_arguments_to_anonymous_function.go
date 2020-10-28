package main
import "fmt"
func update(a *int,t *string){
	*a=*a+5
	*t=*t+" Doe"
	return
}
var(
	area=func(l int,b int)int{
		return l*b 
	}
)
func main(){
	var age=20
	var text="John"
	fmt.Println("Before :",text,age)
	update(&age,&text)
	fmt.Println("After :",text,age)
	fmt.Println()
	fmt.Println("Area of rectangle :",area(20,30))
	func(x int){
		fmt.Println("x^3 :",x*x*x)
		}(20)
}