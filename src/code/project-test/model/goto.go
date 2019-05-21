// goto 可以无条件地转移到程序中指定地行


package model
import (
	"fmt"
)
func GotoTest() {
	var n int = 30
	fmt.Println("ok1")
	fmt.Println("ok2")
	fmt.Println("ok3")
	if n > 20 {
		goto label0
	}
	fmt.Println("ok4")
	fmt.Println("ok5")
	fmt.Println("ok6")
	label0:
	fmt.Println("ok7")
	
}