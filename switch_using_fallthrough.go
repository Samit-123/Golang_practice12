package main
import (
"fmt"
"time"
) 
func main(){
	today := time.Now()
	switch today.Day(){
	case 5:
		fmt.Println("Clean your house.")
		fallthrough
	case 10:
		fmt.Println("Buy some food.")
		fallthrough
	case 15:
	     fmt.Println("Party tonight.")
	     fallthrough
	case 25:
	     fmt.Println("Visit a docter.")
	default:
		fmt.Println("No information available that day.")

	}
}