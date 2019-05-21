
package model


// 闭包就是一个函数 和与其相关的引用环境 组成的一个整体（实体）

import (
	"fmt"
)

func AddUpper() func (int) int {
	var n int = 10
	var str = "hello"
	return func (x int) int {
		n = n + x
		str += string(x)
		fmt.Println("str=", str)
		return n
	}
	// 返回一个匿名函数，但是这个匿名函数引用到函数外的 n 
	// 因此这个匿名函数就和 n 形成一个整体，构成闭包

	// 可以这样理解
	// 闭包是类，函数是方法，n是字段

	// 要搞清楚闭包的关键，就是要分析出 返回的函数它使用（引用）到那些变量
	// 因为函数和它引用到的变量共同构成闭包
}

func AddUpper0() {
	f := AddUpper()
	fmt.Println(f(1))  // 11
	fmt.Println(f(2))  // 13
	fmt.Println(f(3))  // 16

	// 当我们反复调用f函数时， 因为n 是初始化一次，因此每调用一次就进行累计
	
}