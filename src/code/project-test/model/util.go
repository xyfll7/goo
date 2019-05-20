package model

import (
	"fmt"
)

var Hello string = "小洋粉"

func WeekDay()  {
	var days int = 97
	var week int = days / 7
	var day int = days % 7
	fmt.Printf("%d个星期零%d天\n", week, day)
}

func Sheshi() {
	var huashi float32 = 134.2
	var sheshi float32 = 5.0 / 9 * (huashi - 100)
	fmt.Printf("对应的摄氏温度=%v\n", sheshi)
}