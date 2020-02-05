package main

import "fmt"

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x //1 ，返回值赋值=5 2，defer x++ 不改变返回值3，ret指令
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x) //改变的是副本
	return 5
}
func f5() (x int) {
	defer func(x *int) {
		*x++
	}(&x) //改变的是副本
	return 5
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	//fmt.Println(f1()) //5
	//fmt.Println(f2()) //6
	//fmt.Println(f3()) //5
	//fmt.Println(f4()) //6
	//fmt.Println(f5()) //6
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}

//x := 1
//y := 2
//calc("A", x, y) calc("A", 1,2) -> "A" 1 2 3
//x = 10
//calc("B", x, y) calc("B",10,2) -> "B" 10 2 12
//y = 20
//defer calc("BB", x, calc("B", x, y)) calc("BB",10,12) -> "BB" 10 12 22
//defer calc("AA", x, calc("A", x, y)) calc("AA",1,3) -> "AA" 1 3 4
