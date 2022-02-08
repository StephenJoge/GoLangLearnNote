package main

import (
	"commCalculate"
	"fmt"
	"math"
	"net/http"
	"unsafe"
)

/*
	Go 支持的基本类型：
	bool
	数字类型
	int8, int16, int32, int64, int
	uint8, uint16, uint32, uint64, uint
	float32, float64
	complex64, complex128
	byte
	rune
	string
*/
//byte 是 uint8 的别名。rune 是 int32 的别名。
//Go 有着非常严格的强类型特征。Go 没有自动类型提升或类型转换
func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//使用Form方法获取入参值，需要使用ParseForm方法进行表单内容初始化，否则无法使用
		r.ParseForm()
		usernames := r.Form["username"]
		fmt.Fprint(w, usernames)
	} else if r.Method == "POST" {
		//使用FormValue方法，不需要使用ParserForm方法，默认函数内部实现，取入参同名函数的首个值
		username := r.FormValue("username")
		fmt.Fprint(w, username)
	}
}

func main() {

	var a int = 89
	b := 95
	fmt.Println("value of a is", a, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))   // a 的类型和大小
	fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b)) // b 的类型和大小

	var c uint = 10
	fmt.Printf("\ntype of c is %T, size of c is %d", c, unsafe.Sizeof(c)) // c 的类型和大小

	d := 10.256
	fmt.Printf("\ntype of d is %T, size of d is %d", d, unsafe.Sizeof(d)) // d 的类型和大小 float64 是浮点数的默认类型

	c1 := complex(5, 7)
	c2 := 5 + 6i
	fmt.Println("value of c1 is", c1, "and c2 is", c2)
	fmt.Printf("type of c1 is %T,type of c2 is %T ,size of c1 is %d,size of c2 is %d", c1, c2, unsafe.Sizeof(c1), unsafe.Sizeof(c2)) //复数类型

	fmt.Println("\n", math.Sqrt(4))

	area, length := commCalculate.Calculate(12.5, 15)
	fmt.Printf("The area is %v,and the length is %v", area, length)
	// _ 在 Go 中被用作空白符，可以用作表示任何类型的任何值。
	area, _ = commCalculate.Calculate(6, 8)
	fmt.Printf("\nThe area is %v", area)
	area, _ = commCalculate.Calculate(6, 5)
	fmt.Printf("\nThe area is %v", area)

	//数组声明
	var arr1 [3]int
	arr1[0] = 11
	arr1[1] = 12
	fmt.Println("\n数组1：", arr1)

	//数组初始化
	arr2 := [3]int{13, 14}
	fmt.Println("数组2：", arr2)

	//可以忽略声明数组的长度，并用 ... 代替，让编译器为你自动计算长度
	arr3 := [...]int{17, 18, 19, 20}
	fmt.Println("数组3的长度为", len(arr3))

	//数组作为参数传递给函数时，它们是按值传递，而原始数组保持不变
	var arr4 [5]int
	fmt.Println("初始arr：", arr4)
	commCalculate.ChangeLocal(arr4)
	fmt.Println("结果arr：", arr4)

	//range数组遍历
	for i, v := range arr3 {
		fmt.Printf("\nindex is %v,value is %v", i, v)
	}

	for i := range arr2 {
		fmt.Printf("\nindex is %v", i)
	}

	var map1 = make(map[string]int)
	fmt.Printf("\ntype is %T\n", map1)
	map1["A"] = 1
	map1["B"] = 2
	fmt.Println(map1)

	var map2 map[string]int
	fmt.Println(map2)
	if map2 = make(map[string]int); map2 == nil { //此处不会输出，因为条件判断前已进行初始化
		fmt.Println("this map is not initialized!")
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
