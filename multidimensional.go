package main 
import "fmt"
func main() {
       //var n1=[]int{1,23,4,5} 
	   var a=[5][2]int{{1,2},{3,4},{5,6},{7,8},{9,10}}
       //var n[10]int
       var i,j int
       //var a[5][2]int

       for i=0;i<5;i++{
       	for j=0;j<2;j++{
       		fmt.Printf("a[%d][%d]=%d\n",i,j,a[i][j])
       	}
       }

} 
