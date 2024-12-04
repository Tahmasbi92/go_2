package main

import "fmt"

func main() {
	originalNumbers := []int{2, 3, 5, 11, 13, 1, 12, 43, 65, 666, 192}
	//   0  1  2   3   4  5   6   7   8    9   10

	// COPY :
	numbers_Of_Original_Numbers := []int{111}
	copy(originalNumbers[3:], numbers_Of_Original_Numbers)

	fmt.Println(originalNumbers)
	fmt.Println(numbers_Of_Original_Numbers)

}
