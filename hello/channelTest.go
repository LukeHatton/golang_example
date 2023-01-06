// 测试golang中的channel
package main

import "fmt"

// 计算加和，结果放到一个channel中
// 之所以要通过channel返回结果，是因为这样可以使用goroutine来进行并行计算
func sum(arr []int, c chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	// 完成计算后，将结果发送给channel
	c <- sum
}

func main2() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{6, 7, 8, 9, 10}
	c := make(chan int)
	defer close(c)
	// 根据函数调用，channel c中会收到两个值，分别是两次函数调用的结果
	go sum(arr1, c)
	go sum(arr2, c)
	// 因为channel是LIFO的结构，因此会先得到后调用的函数的返回值
	x := <-c
	y := <-c
	fmt.Printf("两个计算结果分别为: %v, %v\n", x, y)
}
