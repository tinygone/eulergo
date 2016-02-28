package eulergo

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Problem13() {
	fmt.Println("Project Euler Problem13")
	//numbers := readFile13()
	fmt.Println("maxMul is " + strconv.FormatInt(15, 10))
}

func readFile13() (list []int64) {
	fileName := "C:/Users/tinygone/git/eulergo/src/problem13.txt"
	inputFile, inputError := os.Open(fileName) //变量指向os.Open打开的文件时生成的文件句柄
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n")
		return
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)

	list2D := make([]string, 0, 0)

	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == nil {
			inputString = strings.TrimSpace(inputString)
			list2D = append(list2D, inputString)

		} else if readerError == io.EOF {
			break
		}
	}
	//fmt.Println(list2D)
	return
}

func addLargeNumber(num1 string, num2 string) (result string) {

	return
}

//string is
func reverseString(num1 string) (numReverse string) {
	var buffer bytes.Buffer
	length := len(num1)
	for i := 0; i < length; i++ {
		piece := num1[length-1-i]
		buffer.WriteByte(piece)

	}
	return buffer.String()
}
