package model

import (
	"fmt"
	"math"
)
func GreaterThan18() {
	var age int
	fmt.Println("请输入你的年龄：")
	fmt.Scanln(&age)

	// 在if中直接定义一个变量
	if age := 23; age > 18 {
		fmt.Println("你的年龄大于18，要对自己负责哦~！")
	}
}

func MathIf(a float64, b float64, c float64) {
	// var a float64 = 3.0
	// var b float64 = 100.0
	// var c float64 = 6.0

	m := b * b - 4 * a * c 

	if m > 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a 
		x2 := (-b - math.Sqrt(m)) / 2 * a 
		fmt.Printf("x1=%v x2=%v", x1, x2)
	} else if m == 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a 
		fmt.Printf("x1=%v", x1)
	} else {
		fmt.Printf("无解...")
	}
}