package main

import (
	"fmt"
)

func main() {
	x := 10
	y := 5

	sum := x + y                  // arithmetic operator
	greater := y > x              // comparison operator
	result := (x > y) || (x > 20) //logical operator

	fmt.Println(sum)
	fmt.Println(greater)
	fmt.Println(result)

}
