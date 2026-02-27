package main

import (
	"fmt"
	"strconv"
)

func numToChar(num int) rune {
	if num <= 9 {
		return rune(96 + num)
	} else if num <= 26 {
		return rune(106 + num - 10)
	}
	return rune(32)
}

func convert(char string) string {
	elem, err := strconv.Atoi(char)
	if err != nil {
		return ""
	}
	return string(numToChar(elem))
}

func main() {
	var str string
	fmt.Scan(&str)

	if len(str) <= 2 {
		for _, elem := range str {
			fmt.Print(convert(string(elem)))
		}
		fmt.Print("\n")
		return
	}

	p1 := 0
	p2 := 2

	for p2 < len(str) {
		if str[p2] == '#' {
			strToConv := string(str[p1]) + string(str[p1+1])

			fmt.Print(convert(string(strToConv)))
			p1 += 2
			p2 += 2
		} else {
			fmt.Print(convert(string(str[p1])))
			p1++
			p2++
		}
	}
	fmt.Println(len(str) - p1)
	if len(str)-p1 == 2 {
		fmt.Println(convert(string(str[p1:])))
	} else {
		fmt.Println(convert(string(str[p1+1:])))
	}

	//fmt.Println(convert(string(str[p1+1:])))

}
