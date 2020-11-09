package main

var localA = "G"

// 结果 : GOG
func main() {
	localN() // 打印全局变量
	localM() // 打印局部变量
	localN() // 打印全局变量
}

func localN() { print(localA) }

func localM() {
	localA := "O"
	print(localA)
}