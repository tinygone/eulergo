package eulergo

import (
	"bufio"
	"fmt"
	"io"
	//"strconv"
	"os"
	"strings"
)

func Problem11() {
	fmt.Println("Project Euler Problem11")
	readFile11()
}

func readFile11() {
	fileName := "C:/Users/tinygone/git/eulergo/src/problem11.txt"
	inputFile, inputError := os.Open(fileName) //变量指向os.Open打开的文件时生成的文件句柄
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n")
		return
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)

	lineCounter := 0

	for {
		inputString, readerError := inputReader.ReadString('\n')
		//inputString, readerError := inputReader.ReadBytes('\n')
		lineCounter++
		fmt.Printf("%d : %s", lineCounter, inputString)

		lineNumberStrList := strings.Split(inputString, " ")
		for _, numStr := range lineNumberStrList {
			

		}
		if readerError == io.EOF {
			return
		}
	}
}
