package main

import "fmt"

func main() {

	for i := 0; i <= 10; i ++ {
		// %b 以二进制输出
		fmt.Printf("num=%3d,  i=%5b,  ^i=%5b\n", i, i, ^i)
	}
}