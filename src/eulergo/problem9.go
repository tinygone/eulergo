package eulergo

import (
	"fmt"
	"strconv"
	//"strings"
)

func Problem9() {
	fmt.Println("Project Euler Problem9")

	var multipleResult int64 = 0
	var sum int64 = 10000000

	var a int64
	var b int64
	var c int64
	var isContinue bool = true
	for a = 2; a < sum; a++ {
		for b = 2; b < sum; b++ {
			c = sum - b - a
			if isPythagoreanTriplet(a, b, c) {
				fmt.Println("find the triplet:")
				isContinue = false
				break

			}
		}
		if !isContinue {
			break
		}
	}
	multipleResult = a * b * c
	fmt.Println("result = " + strconv.FormatInt(multipleResult, 10))

}

func isPythagoreanTriplet(a int64, b int64, c int64) (isP bool) {
	if a*a+b*b == c*c {
		isP = true
	} else {
		isP = false
	}
	return
}
