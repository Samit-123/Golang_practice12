package main

import "fmt"
type First func(int) int
type Second func(int) First

func userdefined(x int)Second{
	return func(y int)First{
		return func(z int)int{
			return x*x+y*y+z*z
		}
	}
}
func main() {
	fmt.Println(userdefined(5)(6)(7))
}