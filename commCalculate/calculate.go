package commCalculate

import "fmt"

// 初始化方法，调用执行
func init() {
	fmt.Println("calculate package initialized")
}

// golang的语法：模块中要导出的函数，必须首字母大写。
func Calculate(height, width float64) (area, length float64) {
	area = height * width
	length = (height + width) * 2
	return area, length
}

func ChangeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("函数内arr：", num)
}
