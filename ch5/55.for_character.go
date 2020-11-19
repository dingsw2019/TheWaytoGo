package main

import (
	"fmt"
	"strings"
)

func forCharacterFunc1(row int) {

	for i := 0; i < row; i ++ {
		for j := 0; j <= i; j ++ {
			print("G")
		}
		println()
	}
}

func forCharacterFunc2(row int) {

	str := "G"
	for i := 0; i < row; i ++ {
		println(str)
		str += "G"
	}
}

func forCharacterFunc3(row int) {

	str := strings.Repeat("G",row)
	for i := 1; i <= row; i ++ {
		fmt.Println(str[:i])
	}
}

func forCharacterFunc4(row int) {

	for i := 1; i <= row; i ++ {
		fmt.Println(strings.Repeat("G",i))
	}
}

func main() {
	row := 25

	forCharacterFunc1(row)
	forCharacterFunc2(row)
	forCharacterFunc3(row)
	forCharacterFunc4(row)
}