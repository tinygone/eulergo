package eulergo

import (
	"fmt"
	"strconv"
)

func Problem3() {
	fmt.Println("Project Euler Problem3")

	var number int64 = 600851475143

	//primeList := make([]int64,10, 10)
	numberFactors := make([]int64, 0, 0)

	var i int64
	tmpnumber := number
	i = 2
	for ; i < tmpnumber/2; i++ {
		if tmpnumber%i == 0 {
			numberFactors = append(numberFactors, i)
			for true {
				if tmpnumber%i == 0 {
					tmpnumber = tmpnumber / i
				} else {
					break
				}
			}
			//fmt.Println(i)
			//fmt.Println(tmpnumber)
		}
	}
	numberFactors = append(numberFactors, tmpnumber)
	fmt.Println(numberFactors)
	maxPrime := findMaxPrime(&numberFactors)

	fmt.Println("max Prime = " + strconv.FormatInt(maxPrime, 10))
}

func isPrime(natural int64) bool {
	var i int64 = 2
	for ; i < natural/2; i++ {
		if natural%i == 0 {
			return false
		}
	}
	return true
}

func findMaxPrime(numbers *[]int64) (maxPrime int64) {
	maxPrime = 0

	for _, tmp := range *numbers {
		if maxPrime < tmp && isPrime(tmp) {
			maxPrime = tmp
		}
	}
	return maxPrime
}
