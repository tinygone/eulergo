package eulergo

import (
	"fmt"
	"strconv"
	//"strings"
)

func Problem6() {
	fmt.Println("Project Euler Problem6")

	sum1 := sumOfSquares(1, 100)
	sum2 := squareOfSum(1, 100)
	diff := sum2 - sum1
	fmt.Println("result = " + strconv.FormatInt(diff, 10))

}

func sumOfSquares(start int64, end int64) (result int64) {
	result = 0
	for i := start; i <= end; i++ {
		result += i * i
	}
	return
}

func squareOfSum(start int64, end int64) (result int64) {
	var sum int64 = 0
	for i := start; i <= end; i++ {
		sum += i
	}
	result = sum * sum
	return
}
