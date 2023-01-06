// 测试golang的反射
package main

import (
	"fmt"
	"reflect"
)

func main() {
	/* 反射三原则 */
	var f float64 = 3.4
	typeOf := reflect.TypeOf(f)
	valueOf := reflect.ValueOf(f)
	fmt.Printf("type of variable f: %v\nvalue of variable f: %v\n", typeOf, valueOf)
	fmt.Println("通过Value获取Type: ", valueOf.Type())

	i := valueOf.Interface()
	fmt.Println(reflect.TypeOf(i))

	s := Student{"ming", 4}
	fmt.Println(reflect.ValueOf(s).Interface())

	fmt.Println("------")
	var x float64 = 3.4
	p := reflect.ValueOf(&x) // Note: take the address of x.
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet()) // settability: false
	elem := p.Elem()                             // Value.Elem()可以返回接口变量或指针类型底层的值
	fmt.Println(elem.CanSet())                   // settability: true
	elem.SetFloat(4.0)
	fmt.Println(elem.Interface())

	/* 利用反射修改结构体字段 */
	fmt.Println("------")
	type T struct {
		A int
		B string
	}
	t := T{23, "tommy"}
	of := reflect.ValueOf(&t)
	s1 := of.Elem()
	fmt.Printf("&t settable? %t | elem settable? %t\n", of.CanSet(), s1.CanSet())
	fmt.Printf("type of variable of: %v | type of variable s1: %v\n", of.Type(), s1.Type())
	typeOfT := s1.Type()
	for i := 0; i < s1.NumField(); i++ {
		field := s1.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, field.Type(), field.Interface())
	}

	fmt.Println("t before modify:", t)
	s1.Field(0).SetInt(17)
	s1.Field(1).SetString("luka")
	fmt.Println("t after modify:", t)
}
