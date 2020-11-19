### for range
> 	for range 模式遍历字符串时, 会根据字符串内容判断字符长度支持遍历 rune(中文)

```go
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for pos, char := range str {
		fmt.Printf("Character on position %d is: %c\n", pos, char)
	}

	fmt.Println()
	fmt.Println()

	str2 := "Chinese: 中文"// 英文字符长度:9 , 中文字符长度:6
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for pos, rune := range str2 {
		fmt.Printf("Character on position %d is: %c\n", pos, rune)
	}
```