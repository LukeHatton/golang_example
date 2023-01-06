// 测试golang的面向对象
package main

import (
	"fmt"
	"strings"
)

// Human 接口：Human
type Human interface {
	Eat()
	Sleep()
}

type Student struct {
	name  string
	grade int
}

// 输出学生信息
func (s Student) String() string {
	return fmt.Sprintf("%d年级的学生%s", s.grade, s.name)
}

func (s Student) Eat() {
	fmt.Printf("%s在吃饭\n", s)
}

func (s Student) Sleep() {
	fmt.Printf("%s在睡觉\n", s)
}

type Worker struct {
	name   string
	salary float64
}

func (w Worker) String() string {
	return strings.ReplaceAll(fmt.Sprintf("薪水为%16.2f的员工%s", w.salary, w.name), " ", "")
}

func (w Worker) Eat() {
	fmt.Printf("%s在吃饭\n", w)
}

func (w Worker) Sleep() {
	fmt.Printf("%s在睡觉\n", w)
}

// 可以通过匿名字段实现继承
type slave struct {
	Worker
}

func main5() {
	humans := []Human{
		Student{name: "ming", grade: 5},
		Worker{name: "qiang", salary: 1000},
	}
	for _, h := range humans {
		switch object := h.(type) {
		// type switch只会判断实际类型，不会在interface类型处就停止匹配
		case Human:
			fmt.Println(object)
		case Student:
			fmt.Println(object)
		case Worker:
			fmt.Println(object)
		default:
			fmt.Println("没有找到对应的匹配类型！")
		}

		h.Eat()
		h.Sleep()
	}

	fmt.Println("------")
	s := slave{Worker{"zheng", 2000}}
	fmt.Println(s)

}
