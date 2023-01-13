// cgo.go 在这个文件中演示如何在golang中调用C函数库，作用就像Java通过JNI调用C函数库一样，只是更加便捷
package main

/*
// 定义一个C语言函数
 #include <stdio.h>

static void sayHello(const char* s){
	puts(s);
}
*/
import "C"

func main() {
	println("hello world with cgo!")
	// 这里没有释放C字符串的内存，在长期运行的程序中会有问题
	C.puts(C.CString("hello world from C!"))
	C.sayHello(C.CString("hello world from custom C function!"))
}
