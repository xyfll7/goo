package model
import (
	"fmt"
)
// go语言没有while 和 do...while语法
// 可以通过for循环来实现其效果


// for循环实现 while 的效果
func WhileTest() {
	i := 0
	for {
		if  i > 10 {
			break
		}
		fmt.Println("Hello, xyf", i)
		i++  // 循环边练迭代
	}	
	fmt.Println("i=", i)
}

// for循环实现do.. while 的效果
func WhileTest0() {
	i := 0
	for {
		fmt.Println("Hello, xyf", i)
		i++  // 循环边练迭代
		if  i > 10 {
			break
		}
	}	
	fmt.Println("i=", i)
}

