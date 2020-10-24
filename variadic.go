package main 
import(
       "fmt"
       "strings"
      ) 
func joinstr(element...string)string{
	return strings.Join(element,"-")
}       
func main(){
	fmt.Println(joinstr())
	fmt.Println(joinstr("Geeks","for","Geeks"))
	fmt.Println(joinstr("A","B","C"))
	element:= []string{"geeks", "FOR", "geeks"} 
	fmt.Println(joinstr(element))
}