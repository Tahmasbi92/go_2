package main

import "fmt"

func main() {
	var number int = 100
	var mypointer *int = &number
	fmt.Println(number) // => ///100
	sum(mypointer)
	fmt.Println(number) ///  => 300
}
func sum(pointer *int) {
	*pointer = *pointer + 200
}
