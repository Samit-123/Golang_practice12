package main 
import "fmt"
func main() {
	var b=[]int{1000,1200,15,5,4}
	avg:=getavg(b,5)
	fmt.Printf("The average of array is %f",avg)
} 
func getavg(arr []int,size int)float32{
	var i,sum int
	var avg float32
	for i=0;i<size;i++{
		sum+=arr[i]
	}
	avg=float32(sum/size)
	return avg
}

