package model
import (
	"fmt"
)
func Cal(n1 float64, n2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	case '%':
		res =float64( int(n1) % int(n2) )
	default:
		fmt.Println("操作符号错误")	
	}
	return res
}

// 函数 调用机制
// 函数 调用过程
// 栈区 基本数据类型 // 逃逸分析
// 堆区 引用数据类型  // 逃逸分析
// 代码区 

func test0(n1 int) {
	n1 = n1 + 1
	fmt.Println("test() n1=", n1)
}

func Test() {
	n1 := 10
	test0(n1)
	fmt.Println("main() n1=", n1)
}

// 在调用一个函数时，会给该函数分配一个新的空间， 编译器会通过自身的处理让这个新的空间和其它的栈空间区分开来
// 在每个函数对应的栈中， 数据空间时独立的， 不会混淆
// 当一个执行完毕后，程序会销毁这个函数对应的栈空间


func GetSumAndSub(n1 int, n2 int) (int ,int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}


// 递归调用

func Recursion(n int) {
	if n > 2 {
		n--
		Recursion(n)
	}
	fmt.Println("n=", n)
}
func Recursion0(n int) {
	if n > 2 {
		n--
		Recursion0(n)
	} else {
		fmt.Println("n=", n)
	}
}
// 执行一个函数时， 就创建一个新的受保护的独立空间（新函数栈）
// 函数的局部变量时独立的，不会相互影响
// 递归必须向退出递归的条件无线逼近， 否则就时无线递归， 
// 当一个函数执行完毕，或遇到return， 就会返回， 遵守谁调用就将结果返回给谁，
// 当函数执行完毕会这返回时， 该函数本身也会被系统销毁


func Fbn(n int) int{
	if(n == 1 || n == 2) {
		return 1
	} else {
		return Fbn(n - 1) + Fbn(n - 2) 
	}
}

func MonkeyEat(n int) int {
	if n > 10 || n < 1 {
		fmt.Println("输入的天数不对")
		return 0 
	}
	if n == 10 {
		return 1
	} else {
		return (MonkeyEat(n +1) + 1) * 2
	}
}

// 注意事项
// 形参列表可以时多个
// 返回值可以是多个
// 形参列表和返回值的列表的数据类型可以是值类型也可以是引用类型
// 函数命名遵循标识符命名规范，首字母不能是数字，首字母大写该函数可以被本包文件和其它包文件使用， 首字母小写只能被本包文件使用
// 函数中的变量是局部的， 函数外不生效
// 基本数据类型和数组默认都是值传递， 即进行值拷贝，在函数内修改不会影响到原来的值
// 如果希望函数内的变量能修改函数外的变量，可以传入变量的地址& 函数内以指针的方式操作变量
// Go函数不支持重载 (支持可变参数、空接口)


func ptr(n1 *int) {
	*n1 = *n1 + 10
	fmt.Println("Ptr() *n1=", *n1)
}

func Ptr() {
	num := 20
	ptr(&num)
	fmt.Println("Ptr() num=", num)
}


// Go中函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量了，通过该变量可以对函数调用




func getSum0(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

func getSum(n1 int, n2 int) int{
	return n1 + n2
}

// 函数既然是一种数据类型，因为在Go中，函数可以作为形参，并且调用
func GetSum1(){
	a := getSum
	fmt.Println(a(44,66))
	res := getSum0(getSum, 10, 30)
	fmt.Println("res= ", res)
}


// 为了简化数据类型定义，Go支持自定义数据类型
// 基本语法： type 自定义数据类型名 数据类型  // 相当于一个别名
type myInt int  // myInt 等价于 int  ,myInt 和 int 是不同类型
type mySum func(int, int) int //mySum 等价于 func (int, int) int

func TypeTest() {
	type myInt int 
	var num1 myInt
	var num2 int
	num1 = 40
	num2 = int(num1) // 需要显示转换
	fmt.Println("num1=", num1, "num2", num2)
}
type myFunType func (int, int) int

func getSum2(funvar myFunType, num1 int, num2 int) int {
	return funvar(num1, num2)
}
// 支持对函数返回值命名

func cal(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2  //这里也不用 var 或者： 了
	sub = n1 - n2 
	return  // 这里就不用再加返回值 sum, sub 了 
}

// 使用 _ 标识符，忽略返回值

// go 支持可变参数
// 0 到多个参数

// 1 到多个参数
func sum22(n1 int, args... int)(sum int)  {
	sum = n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}
// args 是slice 通过 args[index] 可以访问到各个值
// 如果一个函数的形参列表中有可变参数，则可变参数需要放在形参列表最后


// 多个相同类型参数 可以省略一些类型
func sum66(n1, n2 float64) float64 {
	return 1.11
}

// 练习 交换变量的值

func swap(n1 *int, n2 *int) {
	t := *n1
	*n1 = *n2
	*n2 = t
}

func Swap() {
	a := 10
	b := 20
	swap(&a, &b)
	fmt.Printf("a=%v, b=%v", a, b)
}