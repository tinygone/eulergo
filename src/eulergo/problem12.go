package eulergo

import (
	"fmt"
	//"strconv"
	//	"strings"
)

type NatureNumber struct {
	naturalList []int64
	primeList   []int64
	triangle    int64
}

func Problem12() {
	fmt.Println("Project Euler Problem12")

	//naturalList := make([]int64, 0, 0)
	//naturalList = append(naturalList, 0)
	natureNumber := new(NatureNumber)
	natureNumber.naturalList = make([]int64, 0, 0)
	natureNumber.triangle = 0
	for true {
		natureNumber.getNextTriangle()

		factors := natureNumber.getPrimeFactors(natureNumber.triangle)

		//if len(natureNumber.naturalList)%10000 == 0 {
		fmt.Println("----------------------")
		fmt.Println(natureNumber.naturalList)
		fmt.Println(natureNumber.primeList)
		fmt.Println(natureNumber.triangle)
		fmt.Println(factors)
		//}
		if len(factors) > 500 || natureNumber.triangle > 48 {
			break
		}
	}
	fmt.Println(natureNumber.naturalList)
	fmt.Println(natureNumber.primeList)
	fmt.Println(natureNumber.triangle)
}

func (n *NatureNumber) getNextTriangle() {
	var nextNatural int64 = int64(len(n.naturalList) + 1)
	if IsPrime(nextNatural) && !containsNum(n.primeList, nextNatural) {
		n.primeList = append(n.primeList, nextNatural)
	}
	n.naturalList = append(n.naturalList, nextNatural)
	n.triangle = n.triangle + nextNatural
}

func (n *NatureNumber) getPrimeFactors(natural int64) (factors []int64) {
	//factors refer prime factors
	factors = make([]int64, 0, 0)
	if IsPrime(natural) {
		if !containsNum(n.primeList, natural) {
			n.primeList = append(n.primeList, natural)
		}

		appendDistinctList(&factors, int64(1))
		appendDistinctList(&factors, natural)
		return
	}
	naturalLocal := natural
	for _, prime := range n.primeList {
		for true {
			if naturalLocal%prime == 0 {
				naturalLocal = naturalLocal / prime
				factors = append(factors, prime)
				//fmt.Println("naturalLocal=" + strconv.FormatInt(naturalLocal, 10))
			}
			if naturalLocal == natural || naturalLocal%prime != 0 {
				break
			}
		}
		if IsPrime(naturalLocal) && naturalLocal > 1 {
			factors = append(factors, naturalLocal)
			break
		}
	}

	return
}

func containsNum(numList []int64, num int64) (isContain bool) {
	isContain = false
	for _, numTmp := range numList {
		if numTmp == num {
			isContain = true
			return
		}
	}
	return
}

func getAllFactors(primeFactors []int64) (allFactors []int64) {
	var length int = len(primeFactors)
	for i := 1; i <= length; i++ {
		allFactors = Combine(primeFactors, i, length)
	}
	return
}
func Combine(primeFactors []int64, count int, length int) (result []int64) {

	combine_increase(primeFactors, 0, count, count, length, &result)
	return
}

func combine_increase(primeFactors []int64, start int, count int, NUM int, arr_len int, result *[]int64) {
	i := 0
	for i = start; i < arr_len+1-count; i++ {
		(*result)[count-1] = int64(i)
		if count-1 == 0 {
			var j int
			for j = NUM - 1; j >= 0; j-- {
				fmt.Println(*result)
				//printf("%d\t", arr[result[j]])
				//printf("\n")
			}
		} else {
			combine_increase(primeFactors, i+1, count-1, NUM, arr_len, result)
		}
	}
	return
}

func appendDistinctList(numList *[]int64, num int64) {
	if !containsNum(*numList, num) {
		*numList = append(*numList, num)
	}
}
