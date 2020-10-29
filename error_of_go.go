package main
import "fmt"
import "math"
import "errors"
func Sqrt(value float64)(float64,error){
	if (value<0){
		return 0,errors.New("Math: negative number passed to Sqrt")
	}
	return math.Sqrt(value),nil
}
func main(){
	fmt.Printf("First value which is negative :")
	result,err:=Sqrt(-64)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(result)
	}
	fmt.Printf("Second value which is positive :")
	result,err=Sqrt(64)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(result)
	}
}