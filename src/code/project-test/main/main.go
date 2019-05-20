package main  // 表示该 hello.go 文件所在的包是 main
			  // 每个文件归属于一个包

import (
	"fmt"   // 没有使用一个包又不想删除
    "strconv"
    "code/project-test/model"  
) 

func main()  {
	fmt.Println("hello.go")	
	fmt.Println("hello.go")	
	toString()
	fmt.Println(":::::::::::")	
    useStrconv()
    stringToBase()
    ptr()
    fmt.Println(model.Hello)
    model.WeekDay()
    model.Sheshi()
    //model.StudentInfo0()
    //model.GreaterThan18()
    model.MathIf(2.0, 4.0, 2.0)
    model.SwitchTest()
}





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



