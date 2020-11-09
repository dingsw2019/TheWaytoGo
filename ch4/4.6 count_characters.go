package main

import (
	"fmt"
	"unicode/utf8"
)

func main()  {
	str1 := "asSASA ddd dsjkdsjs dk"
	fmt.Println(str1)
	fmt.Printf("the number of bytes in string str1 is %d\n", len(str1))
	fmt.Printf("the number of characters in string str1 is %d\n", utf8.RuneCountInString(str1))

	fmt.Println()

	str2 := "asSASA ddd dsjkdsjsこん dk"
	fmt.Println(str2)
	// 以字节为单位, 长度是 28 ，两个 unicode 占了 4字节
	fmt.Printf("the number of bytes in string str2 is %d\n", len(str2))
	// 以字符为单位，长度是 24，每个字符算一个
	fmt.Printf("the number of characters in string str2 is %d\n", utf8.RuneCountInString(str2))

}
