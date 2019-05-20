## 流程控制

### 基本流程
程序由上到下 逐行顺序执行，中间没有任何判断和跳转

### 举例 注意事项
定义变量 合法的向前引用
``` go
func main() {
    var num1 int = 10
    var num2 int = num1 + 10
    fmt.Println(num2)
}