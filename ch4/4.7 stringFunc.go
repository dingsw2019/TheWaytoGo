/*
	strings 和 strconv 包主要函数演示
 */

package main

import (
	"fmt"
	"strings"
)

func main() {

	/*================================ 查询字符串相关函数 ================================*/

	// 1. 前缀和后缀
	var str string = "This is an example of a string"
	// 字符串str 是否以 th 开头, 返回布尔型
	fmt.Printf("%t\n",strings.HasPrefix(str,"Th")) // true
	fmt.Printf("%t\n",strings.HasPrefix(str,"TH")) // false
	// 字符串str 是否以 ing 结尾, 返回布尔型
	fmt.Printf("%t\n",strings.HasSuffix(str,"ing")) // true
	fmt.Printf("%t\n",strings.HasSuffix(str,"str")) // false


	// 2. 包含关系
	fmt.Printf("%t\n",strings.Contains(str,"an"))   // true
	fmt.Printf("%t\n",strings.Contains(str,"test")) // false
	// 按字符查找是否包含, 字符用单引号
	fmt.Printf("%t\n",strings.ContainsRune(str,'a')) // true
	fmt.Printf("%t\n",strings.ContainsRune(str,99)) // false
	fmt.Printf("%t\n",strings.ContainsRune(str,'中')) // false
	// 多个字符是否都包含在字符串中
	fmt.Printf("%t\n",strings.ContainsAny("chain","b")) // false
	fmt.Printf("%t\n",strings.ContainsAny("chain","c")) // true
	fmt.Printf("%t\n",strings.ContainsAny("chain","c & i")) // true


	// 3. 索引位置
	fmt.Printf("%d\n",strings.Index(str,"an"))   // 8
	fmt.Printf("%d\n",strings.Index(str,"test")) // -1
	fmt.Printf("%d\n",strings.LastIndex(str,"a")) // 22
	// 按字符查找索引位, 字符用单引号
	fmt.Printf("%d\n",strings.IndexRune("chain",rune('c'))) // 0
	fmt.Printf("%d\n",strings.IndexRune("chain",'s')) // -1
	fmt.Printf("%d\n",strings.IndexRune("i'm a 程序猿",'程')) // 6


	// 4. 统计次数
	// strings.Count(s, str string) int
	// 统计 str 在 s 中出现的次数
	fmt.Println(strings.Count("banana","a"))  // 3
	fmt.Println(strings.Count("banana","ba")) // 1
	fmt.Println(strings.Count("banana","c"))  // 0
	fmt.Println(strings.Count("apple banana"," ")) // 1


	// 5. 重复字符串
	// strings.Repeat(s, count int) string
	// 将 s 重复 count 次并返回
	fmt.Println(strings.Repeat("abc",2)) // abcabc


	/*================================ 修改字符串相关函数 ================================*/

	// 6. 字符串替换
	// func Replace(s, old, new string, count int)string
	// 将 s 中的 old 替换成 new，n代表替换次数，-1为全部替换
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 5))  // oinky oinky oinky
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))  // oinky oinky oink
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", -1)) // oinky oinky oinky


	// 7. 字符串大小写
	fmt.Println(strings.ToLower("ABC")) // abc
	fmt.Println(strings.ToUpper("def")) // DEF


	// 8. 修剪字符串
	// 去掉字符串两端的空白字符
	fmt.Println(strings.TrimSpace(" trim space ")) // "trim space"
	// 去掉字符串两端的指定字符串
	fmt.Println(strings.Trim("$trim string$","$")) // "trim string"
	// 去掉字符串左端的指定字符串
	fmt.Println(strings.TrimLeft("## trim left ##","##"))  // " trim left ##"
	// 去掉字符串右端的指定字符串
	fmt.Println(strings.TrimRight("## trim left ##","##")) // "## trim left"


	// 9. 分割字符串，将字符串转成数组
	// 以空白符作为分割符，  \t \n \v \f \r ' ' 都是分割符
	fmt.Printf("%q\n",strings.Fields("Hello 世\n界!\tHe\vl\flo!")) // ["Hello" "世" "界!" "He" "l" "lo!"]
	// 指定字符串为分割符
	fmt.Printf("%q\n",strings.Split("a,b,c",",")) // [a, b, c]
	fmt.Printf("%q\n",strings.Split(" xyz ",""))  // [" " "x" "y" "z" " "]
	fmt.Printf("%q\n",strings.Split("","xyz"))    // [""]


	// 10. 数组转字符串，按指定字符串连接
	arr := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(arr,",")) // foo,bar,baz


}
