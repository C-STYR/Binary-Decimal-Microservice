package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	binaryString := "1001100111"
	fmt.Println(binaryToDecimal(binaryString))
	decimal := 615
	fmt.Println(decimalToBinary(decimal))
	fmt.Println(binaryString == decimalToBinary(decimal))
}

func binaryToDecimal(input string) (num int) {
	num = 0
	var j float64 = 0
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] == 49 {
			num += int(math.Pow(2, j))
		}
		j++
	}
	return
}

func decimalToBinary(num int) (output string) {
	output = ""
	remainders := []int{}
	currentNum := num
	for {
		remainders = append(remainders, currentNum % 2)
		quotient := currentNum / 2
		if quotient == 0 {
			break
		}
		currentNum = quotient
	}
	for i := len(remainders) - 1; i >= 0; i-- {
		output += strconv.Itoa(remainders[i])
	}
	return
}