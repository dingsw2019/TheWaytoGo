package main

var globalA = "G"

// 结果 ： GOO
// 只有一个变量，所有修改都是多全局变量的
func main() {
	globalN()
	globalM()
	globalN()
}

func globalN() {
	print(globalA)
}

func globalM() {
	globalA = "O" // 注意看, 不是声明是赋值
	print(globalA)
}