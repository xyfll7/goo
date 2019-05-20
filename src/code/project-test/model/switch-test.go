package model

import (
	"fmt"
)

func SwitchTest() {
	var key byte
	fmt.Println("请输入一个字符a, b, c, d, e, f, g")
	fmt.Scanf("%c", &key)

	switch test(key) + 1 {
	case 'a':
		fmt.Println("周一，猴子当大王")
	case 'b':
		fmt.Println("周二, 猴子爬雪山")
	case 'c':
		fmt.Println("周三，猴子偷桃")
	default: 
		fmt.Println("输入有误...")	
	}


	var n1 int32 = 20
	//var n2 int64 = 20
	var n3 int32 = 20
	switch n1 {
	// case n2:      //  switch case 匹配数据类型要保持一致
	// 	fmt.Println("ok1")
	case n3, 10 , 5:      //  switch case 匹配数据类型要保持一致
		fmt.Println("ok1")
	// case  5:      //  常量匹配值不能重复
	// 	fmt.Println("ok1")
	default:
		fmt.Println("没有匹配到")
	}
	
	age := 10
	switch {
	case age == 10:
		fmt.Println("age == 10")
	case age == 20:
		fmt.Println("age == 20")
	}

	switch grade := 90; {
	case grade == 10:
		fmt.Println("grade == 10")
	case grade == 90:
		fmt.Println("grade == 90")
	}
	switch grade := 90; {
	case grade == 10:
		fmt.Println("grade == 10")
		fallthrough   // 穿透，继续执行下面代码 默认只能穿透一层
	case grade == 90:
		fmt.Println("grade == 90")
	}

	var x interface{}
	var y = 10.10
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型 %T", i)
	case int:
		fmt.Println("x 是 int 型")
	case float64:
		fmt.Println("x 是 float64")
	case func(int) float64:
		fmt.Println("x 是 func(int) 型")
	case bool, string:
		fmt.Println("x 是 bool 或 string 型")
	default:
		fmt.Println("未知型")
	}

	var month byte
	fmt.Println("输入月份")
	fmt.Scanln(&month)
	switch month {
	case 3, 4, 5:
		fmt.Printf("春天")
	case 6, 7, 8:
		fmt.Println("夏天")
	case 9, 10, 11:
		fmt.Println("秋天")
	case 12, 1, 2:
		fmt.Println("冬天")
	default:
		fmt.Println("输入有误")
	}
}

func test(char byte) byte {
	return char + 1
}