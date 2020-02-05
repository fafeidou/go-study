package main

import (
	"errors"
	"fmt"
)

//类型简写
func intSum(x, y int) int {
	return x + y
}

//可变参数
func intSum2(x ...int) int {
	fmt.Println(x)
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

//返回值
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - 7
	return sum, sub
}

//返回值命名
func calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - 7
	return
}

//函数类型与变量
type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

var c calculation

//高阶函数
//函数作为参数
func calc3(x, y int, op func(int, int) int) int {
	return op(x, y)
}

//函数作为返回值
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别操作符")
		return nil, err
	}
}

func main() {
	count := calc3(10, 20, add)
	fmt.Println(count)

	f, _ := do("+")
	fmt.Printf("%T\n", f)

}
