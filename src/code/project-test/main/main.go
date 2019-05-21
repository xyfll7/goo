package main  // 表示该 hello.go 文件所在的包是 main
			  // 每个文件归属于一个包

import (
	"fmt"   // 没有使用一个包又不想删除
    "strconv"
    _ "code/project-test/model"  
    "code/project-test/utils"  
) 
// 每一个源文件都一个包含一个init函数，
// 该函数会在main函数执行前被GO运行框架调用，
// init函数会在main函数前被调用
// 再该函数中可以完成初始化工作


var age = test()
func test() int{
    fmt.Println("test()")
    return 90
}
func init() {
    fmt.Println("main init...")
}
// 注意事项
// 如果一个文件同时包含全局变量定义，init函数和main 函数， 则执行流程是： 变量定义 > init 函数 > main函数

// main 函数最主要的作用，就是完成一些初始化的工作，
func main()  {
	// fmt.Println("hello.go")	
	// fmt.Println("hello.go")	
	// toString()
	// fmt.Println(":::::::::::")	
    // useStrconv()
    // stringToBase()
    // ptr()
    // fmt.Println(model.Hello)
    // model.WeekDay()
    // model.Sheshi()
    // model.StudentInfo0()
    // model.GreaterThan18()
    // model.MathIf(2.0, 4.0, 2.0)
    // model.SwitchTest()
    // model.WhileTest0()
    // model.WhereIs99()
    // model.Sum20()
    // model.Login()
    // model.ContinueTest()
    // model.Over0()
    // model.GotoTest()
    // var n1 float64 = 1.2
	// var n2 float64 = 2.3
	// var operator byte = '%'
	// result := model.Cal(n1, n2, operator)
    // fmt.Println("result=", result)	
    // result = utils.Cal(n1, n2, operator)
    // fmt.Println("result=", result)
    // model.Test()	
    // res1, res2 := model.GetSumAndSub(3, 2)
    // fmt.Println(res1, res2)
    // res3, _ := model.GetSumAndSub(3, 2)
    // fmt.Println(res3)
    // model.Recursion(40)
    // res := model.Fbn(8)
    // fmt.Println("res =", res)
    // fmt.Println("res =", model.Fbn(4))
    // fmt.Println("res =", model.Fbn(5))
    // fmt.Println("res =", model.Fbn(6))
    // fmt.Println(model.MonkeyEat(1))
    // model.Ptr()
    // model.GetSum1()
    // model.TypeTest()
    fmt.Println("Age=", utils.Age)
    fmt.Println("Age=", utils.Name)
}


// 包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包
// 包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包包
// 包地本质就时创建不同地文件夹 来存放程序文件
// go地每一个文件都属于一个包
// go以包地形式来管理文件和项目目录结构
//
// 包地三大作用
// 1、区分相同名字地函数，变量等标识符
// 2、当程序文件很多时，可以很好的管理项目
// 3、控制函数、变量等访问范围，即作用域

// 包的基本语法
// 打包基本语法
// package util
// 引入包的基本语法
// import "包的路径"

// 包 的注意事项
// 在给一个文件打包时，该包对应一个文件夹，
// 文件的包名通常和文件所在的文件夹名一致
// 包名一般为小写

// 当一个文件要使用其它包函数或变量时，需要首先引入对应的包
// package指令在文件第一行
// 然后时 import指令
// import包时，路径通$GOPATH 的 src 下开始， 不用带src，编译器会自动从src下开始引入
// 为了让其它包的文件，可以访问到本包的函数， 则该函数名的首字母需要大写，（类似于其它语言的public 这样才能跨包访问。比如utils.go
// 在访问其它包函数时，其语法时 包名.函数名 
// 如果包名较长，Go支持给包取别名，取别名后，原来的包名就不能再使用了
// 再同一个包下不能又相同的函数名 也不能有相同的全局变量
// 如果你要编译成一个可执行程序文件，就需要将这个包声明为main，即 package main
// 如果你写的是一个库，包名可以自定义

// 编译  在$GOPATH目录下执行
// go build code/project-test/main
// go build -o bin/my.exe code/project-test/main
// 还会生成一个库文件pkg


func ptr() {
    var i int = 10 
    fmt.Println("i的地址=", &i)

    var ptr *int = &i
    fmt.Printf("ptr=%v\n", ptr)
    fmt.Printf("ptr 的地址=%v\n", &ptr)
    fmt.Printf("ptr 指向的值=%v\n", *ptr)


    var num int = 9
    var ptr0 *int
    ptr0 = &num 
    *ptr0 = 10
    fmt.Println("num=" ,num)
    fmt.Println("*ptr=",*ptr0)
}


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
func useStrconv() {
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
}



func toString() {
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



