package main

import "fmt"

func main() {
	width, height := 20, 10

	for h := 0; h < height; h ++ {
		for w := 0; w < width; w ++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}