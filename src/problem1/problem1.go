package main

import (
	"fmt"
)

func main() {
	fmt.Println("Project Euler Problem1")
	//numbers3_5 := []int{}
	numbers3_5 := make([]int, 0, 1000)
	max_number := 1000

	for a := 2; a < max_number; a++ {
		if a%3 == 0 || a%5 == 0 {
			fmt.Println(a)
			//if there is none of "numbers3_5 =", a error of append(numbers3_5, 1) evaluated but not used will happend
			numbers3_5 = append(numbers3_5, a)
		}
	}
	var sum = 0
	for _, b := range numbers3_5 {
		sum += b
	}
	fmt.Print("sum=")
	fmt.Println(sum)
}
