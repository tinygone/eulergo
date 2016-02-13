package eulergo

import (
	"fmt"
	"strconv"
	//"strings"
)

func Problem5() {
	fmt.Println("Project Euler Problem5")

	numberArray := make([]int64, 0, 0)
	var i int64
	for i = 2; i < 21; i++ {
		if len(numberArray) == 0 {
			numberArray = append(numberArray, i)
		} else {
			var isContain bool = false
			for _, b := range numberArray {
				if i%b == 0 {
					isContain = true
					break
				}
			}
			if isContain {
				
			}else {
				numberArray = append(numberArray, i)
			}
		}
	}
	result := MultipleNumberArray(&numberArray)
	fmt.Println("result = " + strconv.FormatInt(result, 10))

}

func MultipleNatureNumber(start int64, end int64) (result int64) {
	result = start
	var result1 int64
	for i := start + 1; i <= end; i++ {
		result1 = result * i
		fmt.Println(strconv.FormatInt(result1, 10) + " = " + strconv.FormatInt(result, 10) + " * " + strconv.FormatInt(i, 10))
		result = result1
	}
	return result
}

func MultipleNumberArray(numberArray *[]int64) (result int64) {
	result = 1
	var result1 int64
	for _, i := range *numberArray {
		result1 = result * i
		fmt.Println(strconv.FormatInt(result1, 10) + " = " + strconv.FormatInt(result, 10) + " * " + strconv.FormatInt(i, 10))
		result = result1
	}
	return result
}
