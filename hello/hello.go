package main

import (
	"fmt"
	"log"

	"com/example/greetings"
)

func main3() {
	// 设置log使用的前缀
	log.SetPrefix("greetings: ")
	// 设置flag=0，disable时间、源文件、行数的输出
	log.SetFlags(0)

	message, err := greetings.Hello("Tom")
	// 处理错误
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	fmt.Println("------我是分割线😄------")

	names := []string{"Ann", "Mary", "Serin"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
