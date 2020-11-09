package main

var a string

// 结果：GOG
func main() {
	a = "G"
	print(a) // 全局变量
	f1()
}

func f1() {
	a := "O" // 声明语句
	print(a) // 局部变量
	f2()
}

func f2() {
	print(a) // 全局变量
}