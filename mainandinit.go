package main
import (
      "fmt"
      "sort"
      "strings"
      "time"
      )

func main(){
	s:=[]int{45, 78, 123, 10, 76, 2, 567, 5}
	sort.Ints(s)
	fmt.Println("Sorted slice:",s)
	fmt.Println("Index value:",strings.Index("GeeksforGeeks","ks"))
	fmt.Println("Time: ",time.Now().Unix())
}