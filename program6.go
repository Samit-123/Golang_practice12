package main 
import "fmt"
func main(){
	var g string="A"
	switch{
	case g=="A":
		fmt.Println("Good grade")
	case g=="B":
		fmt.Println("ok grade")
	case g=="C":
		fmt.Println("not so good grade")
	case g=="D":
		fmt.Println("Bad grade")		
	default:
		fmt.Println("default")
	}
}