package eulergo

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Problem11() {
	fmt.Println("Project Euler Problem11")
	list2D := readFile11()
	var i int = 0
	var j int = 0
	var maxMul int64 = 0
	for i = 0; i < len(list2D)-3; i++ {
		for j = 0; j < len(list2D[i])-3; j++ {
			mul1 := list2D[i][j] * list2D[i][j+1] * list2D[i][j+2] * list2D[i][j+3]
			mul2 := list2D[i][j] * list2D[i+1][j] * list2D[i+2][j] * list2D[i+3][j]
			mul3 := list2D[i][j] * list2D[i+1][j+1] * list2D[i+2][j+2] * list2D[i+3][j+3]
			mul4 := list2D[i+3][j] * list2D[i+2][j+1] * list2D[i+1][j+2] * list2D[i][j+3]
			maxTmp := getMax(mul1, mul2, mul3, mul4)
			if maxTmp > maxMul {
				maxMul = maxTmp
			}
		}
	}
	fmt.Println("maxMul is " + strconv.FormatInt(maxMul, 10))
}

func readFile11() (list2D [][]int64) {
	fileName := "C:/Users/tinygone/git/eulergo/src/problem11.txt"
	inputFile, inputError := os.Open(fileName) //变量指向os.Open打开的文件时生成的文件句柄
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n")
		return
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)

	lineCounter := 0
	list2D = make([][]int64, 0, 0)

	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == nil {
			lineCounter++
			lineNumbers := make([]int64, 0, 0)
			lineNumberStrList := strings.Split(inputString, " ")
			for _, numStr := range lineNumberStrList {
				numStr = strings.TrimSpace(numStr)
				numTmp, err := strconv.ParseInt(numStr, 10, 64)
				if err == nil {
					lineNumbers = append(lineNumbers, numTmp)
				} else {
					fmt.Println("读取数据有误")
				}

			}
			list2D = append(list2D, lineNumbers)

		} else if readerError == io.EOF {
			break
		}
	}
	//fmt.Println(list2D)
	return
}

func getMax(nums ...int64) (maxSum int64) {
	maxSum = 0
	for _, num := range nums {
		if num > maxSum {
			maxSum = num
		}
	}
	return
}
