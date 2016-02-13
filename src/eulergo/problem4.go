package eulergo

import (
	"fmt"
	"strconv"
	//"strings"
)

func Problem4() {
	fmt.Println("Project Euler Problem4")

	var maxPalindromic int64 = 0
	var a int64 = 100
	var b int64 = 100
	for ; a < 1000; a++ {
		for b = 100; b < 1000; b++ {
			//fmt.Println(a * b)
			if isPalindrome(a*b) && maxPalindromic < a*b {
				maxPalindromic = a * b
				fmt.Println(strconv.FormatInt(a, 10) + "*" + strconv.FormatInt(b, 10))
			}
		}
	}
	fmt.Println("max Prime = " + strconv.FormatInt(maxPalindromic, 10))

}
func isPalindrome(number int64) bool {
	numberStr := strconv.FormatInt(number, 10)
	n := len(numberStr)
	for i := 0; i < n; i++ {
		if numberStr[i] != numberStr[n-i-1] {
			return false
		}
	}
	//fmt.Println(numberStr)
	return true
}
