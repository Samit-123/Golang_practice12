package main
import (
"fmt"
) 
func main(){
	fmt.Println(Calculation("Rectangle",20,30))
	fmt.Println(Calculation("Square",20))
}
func Calculation(str string,y ...int)int{
	area:=1
	for _,value:=range y{
		if str=="Rectangle"{
			area=area*value
		}else if str=="Square"{
			area=value*value
		}
	}
	return area
}