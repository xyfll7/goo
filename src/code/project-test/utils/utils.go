package utils 
import (
	"fmt"
)
func Cal(n1 float64, n2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	case '%':
		res =float64( int(n1) % int(n2) )
	default:
		fmt.Println("操作符号错误")	
	}
	printTest()
	return res
}