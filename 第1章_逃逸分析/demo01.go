package main

import "fmt"

// 返回值
func f1() int {
	t1 := 1 // 不发生逃逸
	return t1
}

// 返回引用,指针
func f2() *int {
	t2 := 3 // 逃逸到堆上
	return &t2
}

//  interface{}动态类型逃逸
func f3(x3 *int) {
	fmt.Println(*x3) // *x3逃逸到堆
}

func main() {
	x1 := f1()
	fmt.Println(x1) // x1逃逸到堆
	x2 := f2()
	fmt.Println(*x2) // *x2逃逸到堆
	f3(x2)
}
