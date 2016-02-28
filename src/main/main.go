package main

import (
	"eulergo"
	"fmt"
)

func main() {
	fmt.Println("start")
	//	eulergo.Problem1()
	//	eulergo.Problem2()
	//	eulergo.Problem3()
	//	eulergo.Problem4()
	//	eulergo.Problem5()
	//	eulergo.Problem6()
	//	eulergo.Problem7()
	//	eulergo.Problem8()
	//	eulergo.Problem9()
	//	eulergo.Problem10()
	//	eulergo.Problem11()
	//eulergo.Problem12()
	//fmt.Println(eulergo.IsPrime(6))
	a := make([]int64, 0, 0)
	a = append(a, 1, 3, 3, 4, 5)
	//e(primeFactors []int64, count int, length int)
	l := eulergo.Combine(a, 1, len(a))
	fmt.Println(l)

}
