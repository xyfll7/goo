// continue 用于结束本次循环，继续执行下一次循环
// 多层循环 也可以用 标签lable跳过某层循环

package model
import (
	"fmt"
)
func ContinueTest() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				continue
			}
			fmt.Println("j=", j)
		}
	}
}

func Jishu() {
	for i := 1; i <= 100; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println("奇数时", i)
	}
}

func Over0() {
	var positiveCount int
	var negativeCount int 
	var num int 
	for {
		fmt.Println("请输入一个整数")
		fmt.Scanln(&num)
		if num == 0 {
			break
		}

		if num > 0 {
			positiveCount++
			continue
		}
		negativeCount++
	}
	fmt.Printf("整数的个数时%v, 负数的个数时%v\n", positiveCount, negativeCount)
}