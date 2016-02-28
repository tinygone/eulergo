package eulergo

import (
	"fmt"
	"strconv"
	//"strings"
)

func Problem10() {
	fmt.Println("Project Euler Problem10")

	var sum int64 = 0

	primeList := make([]int64, 0, 0)
	var i int64 = 2
	for i = 2; i < 2000000; i++ {
		if IsPrime(i) {
			primeList = append(primeList, i)
		}

	}

	for _, t := range primeList {
		sum += t
	}

	fmt.Println("result = " + strconv.FormatInt(sum, 10))

}
