package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello 为一个用户生成问候语
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	// 返回一个随机问候语
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos 为一个用户切片生成问候语
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	// 遍历切片，为每个用户都生成一个问候语，并存储于map中
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil

}

// 初始化此函数要使用的变量
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	// 定义一个切片，包含一组字符串
	formats := []string{
		"Hi, %v, Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// 随机返回切片中的一个元素
	return formats[rand.Intn(len(formats))]
}
