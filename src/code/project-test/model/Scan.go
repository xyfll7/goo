package model

import (
	"fmt"
)

func StudentInfo(){
	var name string
	var age byte
	var sal float32
	var isPass bool
	fmt.Println("请输入姓名：")
	fmt.Scanln(&name)

	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)

	fmt.Println("请输入薪水：")
	fmt.Scanln(&sal)

	fmt.Println("请输入是否通过考试：")
	fmt.Scanln(&isPass)

	fmt.Printf(" 名字是 %v \n 年龄是 %v \n 薪水是 %v \n 是否通过考试 %v \n", name, age, sal, isPass)
}

func StudentInfo0() {
	var name string
	var age byte
	var sal float32
	var isPass bool
	fmt.Println("请输入姓名、年龄、薪水、是否通过考试，使用空格隔开:")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf(" 名字是 %v \n 年龄是 %v \n 薪水是 %v \n 是否通过考试 %v \n", name, age, sal, isPass)
}