package main
import (
      "fmt"
      )
func myfunc(p,q int)(int,int,int){
	return p+q,p*q,p-q
}
func calculatearea(p,q int)(rectangle int,square int){
	rectangle=p*q
	square=p*p
	return
}
func main(){
	var myvar1,myvar2,myvar3=myfunc(4,2)
	fmt.Printf("The result of myvar1 is: %d\n",myvar1)
	fmt.Printf("The result of myvar2 is: %d\n",myvar2)
	fmt.Printf("The result of myvar3 is: %d\n",myvar3)
	fmt.Println()
	var area1,area2=calculatearea(3,2)
	fmt.Printf("Area of rectangle is %d:\n",area1)
	fmt.Printf("Area of square id %d:\n",area2)
}