
package model

import (
	"fmt"
)

// Go 支持匿名函数，
// 如果我们某个函数只是希望使用一次，可以考虑使用匿名函数，
// 匿名函数也可以实现多次调用

// 匿名函数使用方式1
// 定义匿名函数时直接调用

func anonymous() {
	n1 := 30
	n2 := 20
	res0 := func (n1 int, n2 int) int {
		return n1 + n2
	} (n1, n2)
	fmt.Println(res0)

	// 匿名函数使用方式2
	// 将匿名函数赋值给一个变量（函数变量），再通过该变量来调用匿名函数
	// a的数据类型是函数类型
	a := func (n1 int, n2 int) int {
		return n1 + n2
	} 
	res1 := a(10,20)
	fmt.Println(res1)

	res2 := Fun1(2, 5)
	fmt.Println(res2)	
}





// 全局匿名函数
// 如果将匿名函数赋给一个全局变量，那么这个匿名函数，就成为一个全局匿名函数，可以再程序有效

var (
	Fun1 = func (n1 int, n2 int) int {
		return n1 * n2
	}
)