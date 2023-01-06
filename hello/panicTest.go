package main

import (
	"fmt"
)

// 为了模拟panic和recover的使用，在函数g中进行递归，终将抛出一个panic
func g(i int) {
	if i > 3 {
		fmt.Println("函数g panicking!")
		panic(fmt.Sprintf("%v", i))
	} else {
		fmt.Printf("函数g的参数: %v\n", i)
	}

	defer fmt.Println("运行函数g中的defer语句")
	fmt.Println("在这行输出语句后会进行对函数g自身的循环调用")
	g(i + 1)
}

func f() {
	// 注册一个recover函数
	defer func() {
		if r := recover(); r != nil {
			// r的值是panic发生在的函数的入参。如果没有入参，可能会得到nil
			fmt.Printf("在函数f中执行recover!r.type = %T, r.value = %v\n", r, r)
		}
	}()
	fmt.Println("尝试在函数f中调用函数g")
	g(0)
	fmt.Println("从函数g中得到正常的返回值")
}

func main4() {
	f()
	fmt.Println("从函数f中得到正常的返回值")
}
