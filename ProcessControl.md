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
不需要 break 
case 后面可以有多个表达式
