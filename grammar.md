### 变量
``` go

// 默认值
var i int
fmt.Println(i)   // 0

// 类型推导
var num = 10.1
var str = "abc"

// 省略var
name := false

// 多变量声明
var n0, n1, n2
fmt.Println(n0, n1, n2)

var a0, a1, a2 = 1, "xyf", false
fmt.Println(a0, a1, a2)

b0, b1, b2 := 1, "xyf", false
fmt.Println(b0, b1, b2)

// 全局变量
var n1 = 100
var n2 = false
var name = "xyf"

var (
    n1 = 100
    n2 = false
    name = "xyf"
)

func main() {
    fmt.Println(n1, n2, name)
}
```

### 变量 注意事项
同一类型范围内 该区域的值可以不断变化
变量再同意作用域内不能重名
变量 = 变量名+值+数据类型 （变量三要素）

### 变量 数据类型
基本数据类型
    值
        整数 int, int8, int32, int64, uint, uint8, uint16, uint32, uint64（默认）
             byte uint8 的别名
             rune int32 的别名
        浮点 float32, float64
             complex64 complex128
    字符（没有专门的字符型，使用byte保存单个字母字符）
        var a0 byte = 'a'
        var a1 byte = '北'  //overflow溢出
        var c2 int = '北'
        fmt.Println("%c %d", a1, c2) 

        '' 单引号括起来
        '\n' 转义
        UTF-8 编码
        字符本质是一个整数，该字符对应的utf-8编码值
        给变量赋值一个数字，格式化输出 %c
        字符类型可以进行运算，因为背后是整数

        存储： 字符 -> 对应码值 -> 二进制
        读取： 二进制 -> 对应码值 -> 字符
    布尔
        true false 只能取这两个
    字符串
        UTF-8
        字符串一旦赋值，就不可变了，
            var str = "hello"
            str[0] = 'a'  （错）
        "" 双引号，会识别转义字符
        `` 反引号，不识别转义字符，原生字符串输出
        + 字符串拼接
        + 加号保留再上一行



复杂数据类型
    指针 Pointer
    数组
    结构体 struct
    管道 Channel
    函数
    切片 slice （动态数组）
    接口 interface
    map （集合）

### 基本数据类型默认值  零值

``` go
var a int  // 0
var b float32 // 0
var c float64 // 0
var isMarried bool // false
var name string // ""
```

### 基本数据类型转 T(v)
必须显示转换 （不能隐式转换）
T(v) 将v值转换为类型T

``` go
var i int32 = 100
var l0 float32 = float32(i)
var l1 int8 = int8(i)
var l2 int64 = int64(i)  // 低精度 -> 高精度 
```
数据类型转换可以是 表示范围小的 -> 表示范围大的  表示范围大的 -> 表示范围小的
被转换的是变量存储的数据（值），变量本身的数据类型没有变化
int64 -> int8 编译不会报错， 转换的结果按溢出处理
``` go
var num1 int64 = 999999
var num2 int8 = int8(num1)  // 溢出
fmt.Println(num2)  // 63
```

``` go 练习1
func main(){
    var n1 int 32 = 12
    var n2 int64
    var n3 int8

    n2 = n1 +20 // 编译错误
    n3 = n1 +20 // 编译错误
    n2 = int64(n1) +20 // 显式转换
    n3 = int8(n1) +20 // 显式转换
}
```
``` go 练习2
func main(){
    var n1 int32 = 12
    var n2 int8
    var n3 int8

    n4 = int8(n1) + 127 // 编译通过， 结果溢出
    n3 = int8(n1) + 128 // 编译不通过 int8 (-128~ 127)
}
```
### 基本数据类型转 string
fmt.Sprintf("%参数", 表达式)

``` go 练习1
func main() {
    var num1 int = 99
    var num2 float64 = 23.456
    var b bool = true
    var mychar byte = 'h'
    var str string

    str = fmt.Sprintf("%d", num1)
    fmt.Printf("str type %T str=%q\n", str, str)

    str = fmt.Sprintf("%f", num2)
    fmt.Printf("str type %T str=%q\n", str, str)

    str = fmt.Sprintf("%t", b)
    fmt.Printf("str type %T str=%q\n", str, str)

    str = fmt.Sprintf("%c", mychar)
    fmt.Printf("str type %T str=%q\n", str, str)
}
```


strconv包的函数
``` go 练习2 strconv 函数
import (
    "fmt"
    "strconv"
)

func main() {
    var num3 int = 99
    var num4 float64 = 23.456
    var b2 bool = true
    var str string


	// 10进制
    str = strconv.FormatInt(int64(num3), 10)
    fmt.Printf("str type %T str=%q\n", str, str)

	// 'f' 格式   10 小数位保留10位   64 float64
    str = strconv.FormatFloat(num4,'f', 10, 64)
    fmt.Printf("str type %T str=%q\n", str, str)

    str = strconv.FormatBool(b2)
    fmt.Printf("str type %T str=%q\n", str, str)

    // Itoa
    var num5 int64 = 5467
    str = strconv.Itoa(int(num5))
    fmt.Printf("str type %T str=%q\n", str, str)
}
```


ps：
空接口可接受任何类型数据

### string 转基本数据类型

``` go 

func stringToBase() {

    var str string = "true"
    var b bool
    b, _ = strconv.ParseBool(str)
    fmt.Printf("b type %T b=%v\n",b, b)

    var str2 string = "123456789"
    var n1 int64
    var n2 int
    n1, _ = strconv.ParseInt(str2, 10, 64)
    n2 = int(n1)
    fmt.Printf("n1 type %T n1=%v\n",n1, n1)
    fmt.Printf("n2 type %T n2=%v\n",n2, n2)

    var str3 string = "123.456"
    var f1 float64
    f1, _ = strconv.ParseFloat(str3, 64)
    fmt.Printf("f1 type %T f1=%v\n", f1, f1)

    // 确保string 能够转换成一个有效的数据， 失败则会转换成 0 或者 false
    var str4 string = "hello"
    var f2 float64
    f2, _ = strconv.ParseFloat(str4, 64)
    fmt.Printf("f2 type %T f2=%v\n", f2, f2)
}
```

### 指针
基本数据类型 变量存的是值（值类型）

获取变量的地址，用&， 
    var num int
    &num

指针类型，变量存的是一个地址，这个地址指向的空间存的才是值
    var ptr *int = &num  *int 是类型

获取指针所指向的值，使用*  
    var *ptr int
    *ptr 获取指向的值

``` go
func ptr() {
    var i int = 10 
    fmt.Println("i的地址=", &i)

    var ptr *int = &i
    fmt.Printf("ptr=%v\n", ptr)
    fmt.Printf("ptr 的地址=%v\n", &ptr)
    fmt.Printf("ptr 指向的值=%v\n", *ptr)
}
```

### 值类型 引用类型

值类型： 基本数据类型 int系列， float系列， bool， string ，数组，结构体struct
    变量直接存储值 
    内存通常再栈中分配

引用类型： 指针、slice切片、map、管道chan、interface
    通常再堆区分配内存
    ref -> 地址 -> 值 

### 标识符 命名规范

包名 跟所在文件夹保持一致 不要和标准库冲突
驼峰命名法
大写 public   如果变量名、 函数名、常量名 大写 则可以被其它包访问
小写 private  如果                      小写 则只可以在本地访问


### 运算符

``` go
// 只有后 ++ --
i ++ // 只能独立使用

a = i++ // 错误
a = i-- // 错误

if i++ > 0 {   // 错误
    fmt.Println("ok")
}
```

### 关系运算符
==
!=
<
>
<=
>=
### 逻辑运算符

&&
||
！

短路与  前面的判断好了就不判断后面的了
短路或  同上

### 赋值运算符
=
+=
-=
*=
/=
%=

### 位运算符
&
|
^
<<
>>

### 其它运算符
* 指针变量
& 取地址

### 明确不支持三元运算符


### 从键盘输入语句
fmt.Scanln()
fmt.Scanf()
