/*
这个包用来进行各种golang语法测试。为了避免在一个main函数中进行所有测试，
每个文件都会包含一个main函数。在完成一个文件的测试后，将它的main函数重命名，
以免影响之后其他文件的测试
*/
// 比较具有颜色的方块的体积，并输出具有最大体积的方块的颜色
package main

import "fmt"

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type color byte

type box struct {
	height, width, depth float64
	color                color
}

type boxList []box

func (b *box) volume() float64 {
	return b.height * b.width * b.depth
}

func (bl boxList) colorOfBiggestBox() color {
	c := color(WHITE)
	v := 0.0
	for _, b := range bl {
		if b.volume() > v {
			v = b.volume()
			c = b.color
		}
	}
	return c
}

// 在使用slice作为receiver或parameter时要注意，被传递的是slice header的指针的拷贝，其包含的array指针和length是不变的
func (bl boxList) biggestBox() int {
	fmt.Printf("在函数中得到的slice header地址：%p\n", &bl)
	v := 0.0
	index := 0
	// 注意：range函数得到的是对slice底层数组的值的拷贝
	for i, b := range bl {
		if b.volume() > v {
			v = b.volume()
			index = i
		}
	}
	return index
}

func (b *box) paintToAnotherColor(c color) {
	b.color = c
}

// 输出颜色的字符串表示。类比Java，这里应该使用枚举类型
func colorString(c color) string {
	colors := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return colors[c]
}

func (b *box) colorString() string {
	return colorString(b.color)
}

func (b *box) toString() {
	fmt.Printf("the color of this box is %s, volume = %f\n", b.colorString(), b.volume())
}

func main1() {
	boxes := boxList{
		box{height: 10, width: 10, depth: 5, color: BLUE},
		box{height: 10, width: 10, depth: 6, color: WHITE},
		box{height: 10, width: 10, depth: 7, color: YELLOW},
		box{height: 10, width: 10, depth: 8, color: BLUE},
		box{height: 10, width: 10, depth: 9, color: BLACK},
		box{height: 10, width: 10, depth: 10, color: WHITE},
		box{height: 10, width: 10, depth: 11, color: RED},
	}
	fmt.Printf("刚刚初始化后的slice header地址：%p\n", &boxes)
	i := boxes.biggestBox()
	fmt.Printf("slice最大元素索引：%d, slice header地址：%p\n", i, &boxes)
	b := &boxes[i]
	b.color = color(BLUE)
	fmt.Printf("slice中最大box的颜色为%s，体积为%f\n", boxes[i].colorString(), boxes[i].volume())
}
