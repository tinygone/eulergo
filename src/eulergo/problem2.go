package eulergo

import (
	"fmt"
	"strconv"
)

func Problem2() {
	fmt.Println("Project Euler Problem2")

	var upperBound int64 = 4000000
	var sumOfEven int64 = 2
	faciArray := [3]int64{1, 2, 3}

	for true {
		faciArray[0] = faciArray[1]
		faciArray[1] = faciArray[2]
		faciArray[2] = faciArray[0] + faciArray[1]
		if faciArray[2] > upperBound {
			break
		}
		if faciArray[2]%2 == 0 {
			sumOfEven += faciArray[2]
		}
	}

	fmt.Println("sum of even = " + strconv.FormatInt(sumOfEven, 10))
}
