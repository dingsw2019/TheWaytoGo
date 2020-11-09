### 文件名
* 以小写字母组成，如 ``scanner.go``
* 文件名由多个部分组成，使用 ``_`` 相连，如 ``scanner_test.go``


### 关键字

|break|default|func|interface|select|
|-----|-----|---|---|---|
|case |defer | go|map | struct|
|chan |else |goto |package |switch |
|const |fallthrough |if |range |type |
|continue |for |import |return |var |


### 标识符
1. ``有效标识符``，以字母或_开头，后面跟字母或数字
<br>以下为有效标识符
    * X56
    * group1
    * _x23
    * өԑ12 

2. ``无效标识符``，关键字或不符合有效标识符规则的
<br>以下为无效标识符
    * 1ab（以数字开头）
    * case （关键字）
    * a+b （不允许运算符）

3. ``_`` ，是一种特殊标识符，被称为空白标识符。可以用于变量的声明和赋值（任何类型都可以赋值给它）
但任何赋给空白标识符的值都将被抛弃

4. 语句用 ``;`` 分隔，但go不推荐这种写法。一行一句，go编译器自动追加 ``;``


### 包的概念、导入和可见性
一个包（简称pkg）由多个 go文件组成，一个程序由多个包组成，且可执行程序必须包含 main包。

如果一个包进行更新并重新编译，所有引入此包的程序也需要重新编译

``可见性``，包中的变量或函数
* 以 _大写开头_ ，代表外部可见，相当于public。
* 以 _小写开头_ ，代表外部不可见，相当于private

``包的别名``，在引入的包名前设置包的别名，如下
```go
package main
import (
    fm "fmt" // 包的别名
)

func main(){
    fm.Println("hello world") // 包中函数的使用，大写开头可使用
}
```

### 函数

##### main函数
* main() 是程序的入口，init()会在main()之前执行。
* 如果main包没有main(),会引发构建错误 ``undefined: main.main``
* main() 没有入参也没有返回值，如果为其添加了入参或返回，会引发构建错误
 ``func main must have no arguments and no return values results.``


### Go 程序的执行顺序
1. 按 ``main包`` 的 import 中的顺序导入其他包，对每个包执行如下流程
2. 如果该包又导入了其他包，则从第一步开始递归执行，但是每个包只会被导入一次
3. 然后按调用相反的顺序，初始化包中常量和变量，init函数(如果有的话)
    * ``pack1`` 引入 ``pack2`` ，``pack2`` 引入 ``pack3``。
    * 加载到 ``pack3`` 时，``pack3`` 没有引入其他包，开始加载pack3中的常量变量，init函数
    * 然后加载 ``pack2`` 的常量变量、init函数。再然后 ``pack1`` 的变量常量、init函数
4. 完成以上步骤后，``main包``同样执行以上过程，最后调用``main函数``开始执行程序

### 类型转换
**Go不存在隐式转换**，所有转换需要显示说明，格式如下
> 类型 B 的变量 = 类型B (类型 A 的值)
> valueOfTypeB = typeB(valueOfTypeA)

```go
a := 5.0
b := int(a)
```
**类型大的值转换到类型小的变量，会出现截断**，比如 int32 转 int16


### 常量

1. **可用于常量的数据类型：** 布尔、数字(整数、浮点和复数)、字符串
2. 常量的值在编译时必须是确定的
```go
package main

const Ln2 = 0.693147180559945309417232121458\
             176568075500134360255254120680009
const Log2E = 1/Ln2
const Billion = 1e9
const hardEight = (1 << 100) >> 97
```

##### 常量定义方法
* 显式类型定义：``const b string = "abc"``
* 隐式类型定义: ``const b = "abc"``，编译器根据变量的值推断其类型

```go
package main

// 并行赋值
const beef, two, c = "eat", 2, "veg"
const (
    Monday, Tuesday, Wednesday = 1, 2, 3
    Thursday, Friday, Saturday = 4, 5, 6
)
```

### 变量 
1. 变量的声明
* ``var`` 关键字声明变量，``var identifier type``
* 变量命名遵循驼峰规则，例如 numShips ，同时满足可见性规则(首字母小写外部不可见,大写外部可见)
* 变量声明后，自动赋予其零值：int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil。
* ``var 声明`` 建议用于定义 __全局变量__ ，而 ``:= 声明``(短声明)建议用于定义 __局部变量__ (函数中)
* 局部变量声明但不使用，编译时报错。全局变量没有此问题

```go
package main

import "fmt"

var a, b *int     // a,b为int指针类型
var b bool = true // 布尔型
var str string    // 字符串

var (
    num int
    isNumber bool   // 外部不可见变量
    StartDate int64 // 外部可见变量
)

// 短声明多个变量并赋值
// 短声明只能用于局部变量(函数中), 且变量必须被使用
func test() {   
    n1, n2, str1 := 5, 6, "abc"
    fmt.Println(n1, n2, str1)
}
```

> **Go为什么将变量类型放在最后呢？**
> 首先看 C 中定义指针的例子，``int *a, b;`` 
> a是指针，b不是。如果要a、b都是指针，要这样写``int *a, *b``. 而Go的写法 ``var a, b *int`` 在变量很多的情况下写的代码更少
> 其次，我们更需要直到变量的类型，而不是名字。所以放在类型放在后面更容易让人记住

2. 省掉类型的声明，**不建议使用**
- Go编译器很聪明，可以根据变量的值推断其类型(编译时完成)，所以可以直接赋值而省去类型
```go
package main

var a = 15
var (   
    b = false
    str = "Go says hello to the world!"
)


```

但是某些场景不能准确推断，如下
```go
package main
var n1 = 2       // 想给n1 int64类型, 但是编译器根据2 只会给int类型, 所以需要指定类型
var n2 int64 = 2
```
    
### 字符串
创建，字符串在被创建后，长度不能修改，相当于定长数组。

格式，两种字符串格式
* 解释字符串，用双引号包围，其中的转义字符会被替换。
转义字符包括：``\n`` 换行符、``\r`` 回车符、``\t`` tab键、``\u``或``\U`` Unicode字符、``\\`` 反斜杠自身
```go
package main
import "fmt"

// 结果
//This is a raw string
//换行生效
func test() {   
    fmt.Println("This is a raw string \n")
}
```

* 非解释字符串，用 `` 包围，原样输出
```go
package main
import "fmt"

// 结果
//This is a raw string \n
func test() {   
    fmt.Println(`This is a raw string \n` )
}
```

调用，字符串内容为纯字节时，可通过索引来获取内容，索引从 0 开始
**这种方式只对纯ASCII码的字符串有效，还有 获取字符串某个字节的地址是非法行为，如：``&str[i]``**
* 字符串 str 的第 1 个字节：``str[0]``
* 第 i 个字节：``str[i - 1]``
* 最后 1 个字节：``str[len(str) - 1]``

拼接，字符串拼接符 ``+``
```go
// + 必须放在第一行，因为编译器行尾自动补全分号
str := "Beginning of the string " +
    "second part of the string"

s := "hel" + "lo,"
s += "world!"
```
