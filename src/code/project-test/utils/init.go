package utils

var Age int 
var Name string


// 初始化Age 、Name
func init() {
	Age = 100
	Name = "tom~"
}

// 如果 mian.go 和 utils.go里面 都有 全局变量定义、init函数
// 先执行 utils.go里面的 后执行mian.go里面的
//