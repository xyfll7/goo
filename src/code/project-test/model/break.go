package model
import (
	"fmt"
	"math/rand"
	"time"
)
func WhereIs99() {
	count := 0
	for {
		rand.Seed(time.Now().UnixNano())
		// 随机生成 1-100的整数
		n := rand.Intn(100) + 1
		fmt.Println("n=", n)
		count++
		if (n == 99){
			break
		}
	}
	fmt.Printf("生成 99 一共使用了 %v 次", count)
}

// break 注意事项
// break 语句出现在多层嵌套的语句块中时， 可以通过标签指明要终止的时那一层语句块
/**
label0: {
label1: 	{
label2: 		{
					break: label2
				}
			}	
		}
*/

// break 默认会跳出最近的for循环
// break 后面可以指定标签 跳出标签对应的for循环
func lable1()  {
	lable0:
	for i := 0; i < 4; i++ {
		lable2:
		for j := 0; j < 10; j++ {
			if j == 2 {
				break lable0
			} else if j == 3 {
				break lable2
			}
			fmt.Println("j=", j)
		}
	}
}


// 练习
func Sum20()  {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum +=i
		if sum > 20 {
			fmt.Println("当sum>20时，当前数是", i)
			break
		}
	}
}


func Login() {
	var name string
	var pwd string
	var loginChance = 3
	for i := 1; i <= 3; i ++ {
		fmt.Println("请输入用户名")
		fmt.Scanln(&name)
		fmt.Println("请输入密码")
		fmt.Scanln(&pwd)
		if name == "xyf" && pwd == "88" {
			fmt.Println("恭喜你登录成功！")
			break
		} else {
			loginChance--
			fmt.Printf("还有%v次登录机会，请珍惜", loginChance)
		}
	}
	if loginChance == 0 {
		fmt.Println("机会用完，没有登录成功！")
	}
}
