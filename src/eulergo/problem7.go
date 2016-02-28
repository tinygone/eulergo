package eulergo

import (
	"fmt"
	"strconv"
	//"strings"
)

func Problem7() {
	fmt.Println("Project Euler Problem7")

	var primeArray = make([]int64, 0, 0)
	primeArray = append(primeArray, 2)
	var destCount = 10001
	var i int64 = 3
	for true {
		if IsPrime(i) {
			primeArray = append(primeArray, i)
		}

		if len(primeArray)%10 == 0 {
			//fmt.Println(primeArray)
		}

		if len(primeArray) > destCount {
			break
		}
		i++
	}

	fmt.Println("result = " + strconv.FormatInt(i, 10))

}
