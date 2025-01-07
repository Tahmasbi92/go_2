package main

import (
	"fmt"
)

func main() {
	age := 17

	if age >= 18 && age <= 35 {
		fmt.Println("You are in 18 to 35.")
	} else if age <= 0 {
		fmt.Println("You are not born!!!")
	} else if age > 0 && age < 18 {
		fmt.Println("You are a young.")
	} else {
		fmt.Println("You are a old.")
	}
}
