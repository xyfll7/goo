package model
import (
	"fmt"
)
func ForLoop0() {
	for i := 0; i <= 10; i++ {
		fmt.Println("你好", i)
	}	
}

func ForLoop1() {
	i := 0
	for ; i <= 10; i++ {
		
	}	
	fmt.Println(i)
}

func ForLoop2() {
	i := 0
	for i <= 10{
		fmt.Println(i)
		i++ 
	}	
	
}

func ForLoop3() {
	k := 0
	for {
		if k <= 10 {
			fmt.Println(k)	
		} else {
			break
		}
		k++ 
	}	
	
}

func ForLoop4()  {
	str := "hello 北京"
	for index, val :=range str {
		fmt.Printf("index=%d val=%c\n", index, val)
	}
}
// 传统方式 不能遍历中文  （原因是 按照字节遍历，一个汉字在utf-8中占用3个字节）
// 解决方式 讲str 转换成 []rune切片
func ForLoop5()  {
	str := "hello 北京"
	for i := 0; i < len(str); i++ {
		fmt.Printf("val=%c\n", str[i])
	}
}
// 解决方式 讲str 转换成 []rune切片
func ForLoop6()  {
	str := "hello 北京"
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("val=%c\n", str2[i])
	}
}

// 练习
func ForLoop7()  {
	var max uint64 = 100
	var count uint64 = 0
	var sum uint64 = 0
	var i uint64 = 1
	for ; i < max; i++ {
		if i % 9 == 0 {
			count ++
			sum += i
			fmt.Printf("count=%v sum=%v\n", count, sum)
		}
		
	}
}

// 练习
func ForLoop8()  {
	var n int = 60
	for i :=0; i <= n; i++ {
			fmt.Printf("%v + %v = %v\n", i, n - i, n)
	}
}

// 嵌套循环
// 不要超过三层
// 外层孙环次数为M次，内层为N次，则内层循环体实际上需要执行 m*n=mn次

func ForLoop9()  {
	var classNum int = 2
	var stuNum int = 5
	var totalSum float64 = 0.0
	var passCount int = 0
	for i := 1; i <= classNum; i++ {
		sum := 0.0
		for j := 1; j <= stuNum; j++ {
			var score float64
			fmt.Printf("请输入第%d班 第%d个学生的成绩 \n", i, j)
			fmt.Scanln(&score)
			sum += score

			if score >= 60 {
				passCount++
			}
		}
		fmt.Printf("第%d个班的平均分是%v\n", i, sum / float64(stuNum))
		totalSum += sum
	}
	fmt.Printf("各个班级的总成绩%v 所有班级平均分是%v\n", totalSum, totalSum / float64(stuNum))
	fmt.Printf("及格人数为%v\n", passCount)
}


// 打印金字塔

func ForLoop10(){
	totalLevel := 9
	for i := 1; i <= totalLevel; i++ {
		for k := 1; k <= totalLevel - i; k++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j ==2*i-1 || i == totalLevel{
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
			
		}
		fmt.Println()
	}
}

func ForLoop11()  {
	// 打印99乘法表
	var num int = 9
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v \t", j, i, j * i)
		}
		fmt.Println()
	}
}

