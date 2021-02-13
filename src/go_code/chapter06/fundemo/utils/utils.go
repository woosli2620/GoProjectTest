package utils

var Num int = 20

//为了其他文件使用，函数首字母需大写
func Cal(n1 float64,n2 float64,oper byte) (float64) {

	var res float64 = 0
	switch oper {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
	}

	return res
}
