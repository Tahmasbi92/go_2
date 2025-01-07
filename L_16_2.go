package main

import "fmt"

func main() {
	person_scores := map[string]float64{
		"sina": 12.5,
	}
	fmt.Println(person_scores)

	var temperatures map[string]float64 = map[string]float64{
		"New York": 10.5,
	}
	fmt.Println(temperatures)
}
