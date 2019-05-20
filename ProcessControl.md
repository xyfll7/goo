## 流程控制

### 基本流程
程序由上到下 逐行顺序执行，中间没有任何判断和跳转

### 举例 注意事项
定义变量 合法的向前引用
正确：✔
``` go
func main() {
    var num1 int = 10
    var num2 int = num1 + 10
    fmt.Println(num2)
}
```
错误：❌
``` go
func main() {
    var num2 int = num1 + 10
    var num1 int = 10
    fmt.Println(num2)
}
```

### 分支控制
#### 单分支
``` go
if xxx {

}
```
必须有大括号{}

#### 双分支
``` go
if xxx {

} else {

}
```
#### 多分枝
``` go
if xxx {

} else if xxx {

} else if xxx {

} else {

}
```
#### 嵌套分支 
最好控制在3层以内

### switch 分支
1、不需要 break 
2、case/swich 后面可以跟表达式
3、数据类型保持一致
4、case 后面可以有多个表达式
5、default不是必须的
6、case 后面表达式如果是常量值要求不能重复

7、switch 后面可以不带表达式
8、switch 后面可以声明定义变量 分号结束
9、fallthrough   // 穿透，继续执行下面代码
10、 Type Switch 判断某个interface变量中实际指向的变量类型
